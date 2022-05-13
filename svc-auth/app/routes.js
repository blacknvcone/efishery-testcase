'use_strict';

const RegisterController = require("./controller/register");

module.exports = (app) => {
    const registerController = new RegisterController();

    app.post("/register", registerController.RegisterUser);

};