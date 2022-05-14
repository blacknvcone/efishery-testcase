'use_strict';

const jwt = require("jsonwebtoken");

class JWTLibrary {
    constructor() {
        this.secret = process.env.JWT_SECRET;

        this.GenerateToken = this.GenerateToken.bind(this);
        this.VerifyToken = this.VerifyToken.bind(this);
    }

    GenerateToken(user, now = new Date().getTime()) {
        const payload = {
            username: user.username,
            phone: user.phone,
            role: user.role,
            timestamp: now
        };
        return jwt.sign(payload, this.secret, {
            expiresIn: '1h'
        });
    }

    VerifyToken(token) {
        return jwt.verify(token, this.secret);
    }
}

module.exports = JWTLibrary;