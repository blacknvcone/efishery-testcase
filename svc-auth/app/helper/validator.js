`use_strict`;

module.exports = {
    validateBody: function (reqBody, requires, strings) {
        const required = "$_variable is required"

        let errors = {}
        for (let i in requires) {
            if (typeof reqBody[requires[i]] === 'undefined' || (reqBody[requires[i]] === "")) {
                errors[requires[i]] = required.replace('$_variable', requires[i])
            }
        }

        if (Object.getOwnPropertyNames(errors).length > 0) {
            return errors
        }
        return false
    }
}