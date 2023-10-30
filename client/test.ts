import { Get, Write } from "./index";

const Test = async () => {
  try {
    await Write({
      Key: "key",
      Commands: "TSET",
      IfNotExist: false,
      Data: {
        value: "🚀",
        EX: 1000,
      },
    });
  } catch (err) {
    console.log(err);
  }

  try {
    const res = await Get({
      Key: "key",
      Command: "GET",
    });
    console.log("respons: ", res);
  } catch (err) {
    console.log(err);
  }
};

Test();
