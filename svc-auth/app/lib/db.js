'use_strict';

const db = require("diskdb");

module.exports = {
    connectDB: function () {
        db.connect(process.env.REGISTRY_PATH, [process.env.REGISTRY_COLLECTIONS])
    }
}