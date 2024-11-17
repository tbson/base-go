import { atom } from 'jotai';

export const tenantOptionSt = atom({
    loaded: false,
    auth_client: []
});

export const tenantFilterSt = atom((get) => {
    const { auth_client } = get(tenantOptionSt);
    const tenantFilter = auth_client.map((item) => ({
        value: item.value,
        text: item.label
    }));
    return {
        auth_client: tenantFilter
    };
});

export const tenantDictSt = atom((get) => {
    const { auth_client } = get(tenantOptionSt);
    const tenantDict = auth_client.reduce((acc, item) => {
        acc[item.value] = item.label;
        return acc;
    }, {});
    return {
        auth_client: tenantDict
    };
});
