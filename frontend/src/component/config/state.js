import { atom } from 'jotai';

export const configOptionSt = atom({
    loaded: false,
    variable_data_type: []
});

export const variableOptionSt = atom({
    loaded: false,
    data_type: []
});
