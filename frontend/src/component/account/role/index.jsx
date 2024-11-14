import * as React from "react";
import PageHeading from "component/common/page_heading";
import Table from "./table";
import { getMessages } from "./config";

export default function Role() {
    const messages = getMessages();
    return (
        <>
            <PageHeading>
                <>{messages.heading}</>
            </PageHeading>
            <Table />
        </>
    );
}

Role.displayName = "Role";
