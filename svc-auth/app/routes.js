'use_strict';

const RegisterController = require("./controller/register");
const SignInController = require("./controller/sign_in");

module.exports = (app) => {
    const registerController = new RegisterController();
    const signInController = new SignInController();

    app.post("/register", registerController.RegisterUser);
    app.post("/signin", signInController.SingIn);

};