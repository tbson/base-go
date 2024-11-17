class TableUitl {
    static optionToFilter(options) {
        return options.map((option) => ({
            value: option.value,
            text: option.label
        }));
    }
}

export default TableUitl;
