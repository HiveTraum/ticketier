### Architectural cases that must be scaled

1. Ticket routing mechanism. Ticket can be routed by many properties of user. 
   Such as department, city, dialog flow answers, etc. This must be scaled at user level. 
   That means that the user can configure this routing.
   1. Organize property based structure? How to manage in admin UI filling data that we don't have in our db? API Adapters?
   
2. 