ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/"
"/dog/"
"/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.

//New Task
1. Take the previous program in the previous folder and change it so that:
* a template is parsed and served
* you pass data into the template
