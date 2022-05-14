const jwt = require('jsonwebtoken');

module.exports = async (req, res, next) => {
    if (!req.headers["authorization"]) {
        return res.sendError({
            Authorization: "Invalid Token!"
        }, "Invalid Token!", 401)
    }

    const token = req.headers["authorization"].split(" ")
    try {
        let decoded = jwt.verify(token[1], process.env.JWT_SECRET)

        req.user = decoded
        return next() // use next to go next router

    } catch(err) {
        return res.sendError(err, "Invalid Token!", 401)
    }
}
