'use_strict';

const RegisterController = require("./controller/register");
const SignInController = require("./controller/sign_in");
const ProfileController = require("./controller/profile");
const AuthMidlleware = require("../middlewares/auth");

module.exports = (app) => {
    const registerController = new RegisterController();
    const signInController = new SignInController();
    const profileController = new ProfileController();

    app.post("/register", registerController.RegisterUser);
    app.post("/signin", signInController.SingIn);

    //Assign Middleware Auth
    app.use(AuthMidlleware)
    app.get("/profile", profileController.DetailProfile);

};