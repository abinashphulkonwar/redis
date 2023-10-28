"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.Get = void 0;
const node_fetch_1 = __importDefault(require("node-fetch"));
const init_1 = require("./init");
const error_1 = require("./error");
const Get = ({ Key, Command, }) => __awaiter(void 0, void 0, void 0, function* () {
    const config = (0, init_1.GetConfig)();
    const request = yield (0, node_fetch_1.default)(`${config.url}/api/query/GET?command=${Command} ${Key}`, {
        method: "GET",
        headers: {
            "content-type": "application/json",
        },
    });
    if (request.ok) {
        const res = (yield request.json());
        console.log(res);
        return res;
    }
    const res = (yield request.json());
    throw new error_1.RedisClientError(res);
});
exports.Get = Get;
(0, exports.Get)({
    Key: "key",
    Command: "GET",
});
