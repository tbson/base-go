import { t } from 'ttag';
import RequestUtil from 'service/helper/request_util';

const urlMap = {
    base: {
        prefix: 'account/user',
        endpoints: {
            crud: '',
            option: 'option'
        }
    }
};
export const urls = RequestUtil.prefixMapValues(urlMap.base);
export const TOGGLE_DIALOG_EVENT = 'TOGGLE_USER_DIALOG';
export const PEM_GROUP = 'cruduser';
const headingTxt = t`User`;
const name = headingTxt.toLowerCase();
export const getMessages = () => ({
    heading: headingTxt,
    deleteOne: t`Do you want to remote this ${name}?`,
    deleteMultiple: t`Do you want to remote these ${name}?`
});

export const getLabels = () => ({
    external_id: t`External ID`,
    sub: t`Sub`,
    email: t`Email`,
    mobile: t`Mobile`,
    first_name: t`First name`,
    last_name: t`Last name`,
    avatar: t`Avatar`,
    admin: t`Admin`,
    locked: t`Locked`,
    locked_reason: t`Locked reason`,
    role_ids: t`Roles`
});
