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
Object.defineProperty(exports, "__esModule", { value: true });
const index_1 = require("./index");
const Test = () => __awaiter(void 0, void 0, void 0, function* () {
    try {
        yield (0, index_1.Write)({
            Key: "key",
            Commands: "TSET",
            IfNotExist: false,
            Data: {
                value: "🚀",
                EX: 1000,
            },
        });
    }
    catch (err) {
        console.log(err);
    }
    try {
        const res = yield (0, index_1.Get)({
            Key: "key",
            Command: "GET",
        });
        console.log("respons: ", res);
    }
    catch (err) {
        console.log(err);
    }
});
Test();
