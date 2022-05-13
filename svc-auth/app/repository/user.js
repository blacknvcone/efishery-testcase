'use_strict';

const db = require('diskdb');

class UserRepository {
    constructor() {
        this.CreateUser = this.CreateUser.bind(this);
        this.GetByUsername = this.GetByUsername.bind(this);
    }

    CreateUser(user) {
        return db.users.save(user);
    }

    GetByUsername(username) {
        const usr = db.users.find({ username: username });
        return usr.length <= 0;
    }
}

module.exports = UserRepository;