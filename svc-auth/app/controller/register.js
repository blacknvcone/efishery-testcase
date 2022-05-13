`use_strict`;

const UserRepository = require("../repository/user");
const validator = require("../helper/validator");
const randomstring = require("randomstring");

class RegisterController {

    constructor() {
        this.RegisterUser = this.RegisterUser.bind(this);
        this.UserRepository = new UserRepository();
    }

    async RegisterUser(req, res) {

        const required = [
            'username', "name", "phone", "role"
        ]
        let err = validator.validateBody(req.body, required)

        if (err) {
            return res.sendError(err)
        }

        //Validate Username
        if (this.UserRepository.GetByUsername(req.body.username)) {

            req.body.password = randomstring.generate(4);
            //// register user
            let addUser = await this.UserRepository.CreateUser(req.body)
            if (addUser) {
                return res.sendSuccess(addUser, "User created!", 200);
            } else {
                return res.sendError(null, "Failed saved data!", 406);
            }

        } else {
            return res.sendError(null, "Username was already used!", 400);
        }
    }
}

module.exports = RegisterController;