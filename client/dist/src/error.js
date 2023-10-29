"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RedisClientError = void 0;
class RedisClientError extends Error {
  constructor({ status, data, message }) {
    super(message);
    Object.setPrototypeOf(this, RedisClientError.prototype);

    this.status = status;
    this.data = data;
    this.message = message;
  }
  serializeErrors() {
    return {
      status: this.status,
      data: this.data,
      message: this.message,
    };
  }
}
exports.RedisClientError = RedisClientError;
