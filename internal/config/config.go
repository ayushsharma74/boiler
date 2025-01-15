package config

// RepositoryInfo defines the details of a boilerplate repository
type RepositoryInfo struct {
    URL         string
    Description string
}

// Define available boilerplate repositories
var Repositories = map[string]RepositoryInfo{
    "basic-express": {
        URL:         "https://github.com/your-username/basic-express-boilerplate", // Replace with an actual repo
        Description: "Basic Express.js project with minimal setup",
    },
    "express-with-ejs": {
        URL:         "https://github.com/your-username/express-ejs-boilerplate", // Replace with an actual repo
        Description: "Express.js with EJS template engine",
    },
    "express-with-hbs": {
        URL:         "https://github.com/your-username/express-hbs-boilerplate", // Replace with an actual repo
        Description: "Express.js with Handlebars template engine",
    },
    // Add more options as needed
}