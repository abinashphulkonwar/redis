"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.GetConfig = exports.init = void 0;
const config = {
    url: "http://localhost:3000",
    token: "",
};
const init = ({ url, token }) => {
    config.url = url;
    config.token = token;
};
exports.init = init;
const GetConfig = () => {
    return config;
};
exports.GetConfig = GetConfig;
