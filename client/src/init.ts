const config = {
  url: "http://localhost:3000",
  token: "",
};

export const init = ({ url, token }: { url: string; token: string }) => {
  config.url = url;
  config.token = token;
};

export const GetConfig = () => {
  return config;
};
