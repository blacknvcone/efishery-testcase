`use_strict`;

const UserRepository = require("../repository/user");
const Validator = require("../helper/validator");
const JWTLibrary = require("../lib/jwt");

class SignInController {

    constructor() {
        this.SingIn = this.SingIn.bind(this);
        this.UserRepository = new UserRepository();
        this.JWTLibrary = new JWTLibrary();
    }

    async SingIn(req, res) {
        const required = [
            "phone", "password"
        ]

        let err = Validator.validateBody(req.body, required);

        if (err) {
            return res.sendError(err)
        }

        //Validate Data
        let resval = await this.UserRepository.GetByPhonePass(req.body.phone, req.body.password);
        if (resval.length > 0) {

            let now = new Date().getTime();
            let jwt = this.JWTLibrary.GenerateToken(resval, now);

            return res.sendSuccess({
                "token": jwt
            }, "User Authenticated!");
        } else {
            return res.sendError(null, "Wrong phone or password !", 401);
        }

    }
}

module.exports = SignInController;