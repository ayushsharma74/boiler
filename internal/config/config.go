package config

// RepositoryInfo defines the details of a boilerplate repository
type RepositoryInfo struct {
    URL         string
    Description string
}

// Define available boilerplate repositories
var Repositories = map[string]RepositoryInfo{
    "basic-express": {
        URL:         "https://github.com/ayushsharma74/basic-express-boilerplate", // Actual Repo Did not exist 
        Description: "Basic Express.js project with minimal setup",
    },
    "express-with-ejs": {
        URL:         "https://github.com/ayushsharma74/express-ejs-boilerplate", // Actual Repo Did not exist 
        Description: "Express.js with EJS template engine",
    },
    "express-with-hbs": {
        URL:         "https://github.com/ayushsharma74/express-hbs-boilerplate", // Actual Repo Did not exist 
        Description: "Express.js with Handlebars template engine",
    },
    // Add more options as needed
}