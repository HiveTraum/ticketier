import * as React from 'react';
import {useContext} from 'react';
import {AppBar, Button, createStyles, makeStyles, Toolbar, Typography} from "@material-ui/core";
import {Link} from "react-router-dom";
import {UserContext} from "../storages/UserStorage";

const useStyles = makeStyles(() =>
    createStyles({
        title: {
            flexGrow: 1,
        },
        button: {
            color: "white"
        },
        link: {
            textDecoration: "none",
        }
    }),
);


const Navigation: React.FunctionComponent = () => {
    const classes = useStyles();
    const {user, signIn, signOut} = useContext(UserContext);
    console.log("navigation render");

    return <AppBar position="static">
        <Toolbar>
            <Typography variant="h6" className={classes.title}>
                <Link to="/" className={classes.link}><Button className={classes.button}>Kiosk</Button></Link>
            </Typography>
            {user && <Typography variant="h6" className={classes.title}>
                {user.fullName}
            </Typography>}

            <Button className={classes.button} onClick={user ? signOut : signIn}>{user ? "Выйти" : "Войти"}</Button>
        </Toolbar>
    </AppBar>
}

export default Navigation;