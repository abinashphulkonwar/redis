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
exports.Write = void 0;
const node_fetch_1 = __importDefault(require("node-fetch"));
const init_1 = require("./init");
const error_1 = require("./error");
const Write = (event) => __awaiter(void 0, void 0, void 0, function* () {
    const config = (0, init_1.GetConfig)();
    const request = yield (0, node_fetch_1.default)(`${config.url}/api/write/add`, {
        method: "POST",
        headers: {
            "content-type": "application/json",
        },
        body: JSON.stringify(event),
    });
    if (request.ok) {
        const res = (yield request.json());
        console.log(res);
        return res;
    }
    const res = (yield request.json());
    console.log(res);
    throw new error_1.RedisClientError(res);
});
exports.Write = Write;
