'use_strict';

class ProfileController {

    constructor() {
        this.DetailProfile = this.DetailProfile.bind(this);
    }

    async DetailProfile(req, res) {
        //Returning embedded request data from auth middleware
        return res.sendSuccess(req.user, "OK", 200);
    }
}

module.exports = ProfileController;