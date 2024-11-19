import { atom } from 'jotai';
import TableUtil from 'service/helper/table_util';

export const roleOptionSt = atom({
    loaded: false,
    auth_client: []
});

export const roleFilterSt = atom((get) => {
    const { auth_client } = get(roleOptionSt);
    return {
        auth_client: TableUtil.optionToFilter(auth_client)
    };
});

export const roleDictSt = atom((get) => {
    const { auth_client } = get(roleOptionSt);
    const roleDict = auth_client.reduce((acc, item) => {
        acc[item.value] = item.label;
        return acc;
    }, {});
    return {
        auth_client: roleDict
    };
});
