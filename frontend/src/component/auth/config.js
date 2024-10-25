import RequestUtil from "service/helper/request_util";

const urlMap = {
    base: {
        prefix: "account/auth/sso",
        endpoints: {
            login: "login",
        }
    }
};

export const urls = RequestUtil.prefixMapValues(urlMap.base);

const headingTxt = "Hồ sơ";
export const messages = {
    heading: headingTxt
};
