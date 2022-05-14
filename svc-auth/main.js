require("express-group-routes");
require("dotenv").config();

const createError = require("http-errors");
const express = require("express");
const cookieParser = require("cookie-parser");
const cors = require("cors");
const logger = require("morgan");
const port = process.env.PORT;

//Init Route,DB, and Middleware
const apiRouter = require("./app/routes");
const generalMiddleware = require("./middlewares/general");
const db = require("./app/lib/db");

db.connectDB();

const app = express();

// enable cors
app.use(
    cors({
        exposedHeaders: ["Content-Disposition"],
    })
);

app.use(logger('dev'));
app.use(express.json({ limit: "100mb" }));
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

//Assign general response formatter middleware
app.use(generalMiddleware);

// Swagger Settings
const swaggerUi = require('swagger-ui-express');
const openApiDocumentation = require('./swagger.json');
app.use('/swagger', swaggerUi.serve, swaggerUi.setup(openApiDocumentation));

//Registering All Service Route
apiRouter(app);

//Handling error for unregistered endpoint
app.use(function (req, res, next) {
    next(createError(404));
});

// error controller
app.use(function (err, req, res, next) {
    res.sendError(err, err.message, err.status);
});

//server listening
app.listen(port, () => {
    console.log(`Server is running at port ${port}`);
});

module.exports = app;