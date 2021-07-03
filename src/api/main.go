package main

import (
	"github.com/minio/minio-go"
	"src/file"
	"src/postgresql"
	"src/subject"
	"src/subject_field"
	"src/system"
	"src/ticket"
	"src/ticket_answer"
	"src/ticket_attachment"
)

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	env, err := system.NewEnvironment()
	panicIfError(err)

	postgresqlDB, err := postgresql.New(env.PostgreSQLURI)
	panicIfError(err)

	minioClient, err := minio.New(env.MinioEndpoint, env.MinioUsername, env.MinioPassword, env.MinioUseSSL)
	panicIfError(err)

	ticketEntity := ticket.NewTicketEntity()
	ticketAnswerEntity := ticket_answer.NewTicketAnswerEntity()
	ticketAttachmentEntity := ticket_attachment.NewTicketAttachmentEntity()
	subjectEntity := subject.NewSubjectEntity()
	subjectFieldEntity := subject_field.NewSubjectFieldEntity()

	ticketRepository := ticket.NewTicketPostgreSQLRepository(postgresqlDB)
	ticketAnswerRepository := ticket_answer.NewTicketAnswerRepositoryPostgreSQL(postgresqlDB)
	ticketAttachmentRepository := ticket_attachment.NewTicketAttachmentRepositoryPostgreSQL(postgresqlDB)
	subjectRepository := subject.NewSubjectPostgreSQLRepository(postgresqlDB)
	subjectFieldRepository := subject_field.NewSubjectFieldPostgreSQLRepository(postgresqlDB)
	ticketAttachmentFileRepository, err := file.NewFileMinioRepository(minioClient, "attachments")
	panicIfError(err)

	ticketService := ticket.NewTicketService(ticketEntity, ticketRepository, ticketAnswerEntity, ticketAnswerRepository, ticketAttachmentEntity, ticketAttachmentRepository, ticketAttachmentFileRepository, subjectFieldRepository, subjectRepository)
	subjectService := subject.NewSubjectService(subjectEntity, subjectRepository, subjectFieldEntity, subjectFieldRepository)

	err = InitREST(ticketService, subjectService)
	panicIfError(err)
}
