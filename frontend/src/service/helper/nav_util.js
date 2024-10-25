import StorageUtil from "service/helper/storage_util";
import RequestUtil from "service/helper/request_util";

export default class NavUtil {
    /**
     * navigateTo.
     *
     * @param {Navigate} navigate
     */
    static navigateTo(navigate) {
        return (url = "/") => {
            navigate(url);
        };
    }

    /**
     * logout.
     *
     * @param {Navigate} navigate
     */
    static logout() {
        return () => {
            const baseUrl = RequestUtil.getApiBaseUrl();
            const tenantUid = StorageUtil.getTenantUid();
            const logoutUrl = `${baseUrl}account/auth/sso/logout/${tenantUid}`;
            StorageUtil.removeStorage("userInfo");
            StorageUtil.removeStorage("tenantUid");
            StorageUtil.removeStorage("locale");
            window.location.href = logoutUrl;
        };
    }

    /**
     * cleanAndMoveToLoginPage.
     *
     * @param {Navigate} navigate
     * @returns {void}
     */
    static cleanAndMoveToLoginPage(navigate) {
        const currentUrl = window.location.href.split("#")[1];
        StorageUtil.removeStorage("userInfo");
        StorageUtil.removeStorage("tenantUid");
        StorageUtil.removeStorage("locale");
        let loginUrl = "/login";
        if (currentUrl) {
            loginUrl = `${loginUrl}?next=${currentUrl}`;
        }
        if (navigate) {
            NavUtil.navigateTo(navigate)(loginUrl);
        } else {
            window.location.href = `/#${loginUrl}`;
        }
    }
}
