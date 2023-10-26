export interface ErrorInterface {
  status: string;
  data: null;
  message: string;
}
export class RedisClientError extends Error {
  status: string;
  data: null;
  message: string;
  constructor({ status, data, message }: ErrorInterface) {
    super(message);
    this.status = status;
    this.data = data;
    this.message = message;
    Object.setPrototypeOf(this, RedisClientError.prototype);
  }
  serializeErrors() {
    return {
      status: this.status,
      data: this.data,
      message: this.message,
    };
  }
}
