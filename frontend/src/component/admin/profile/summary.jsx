import * as React from "react";
import { t } from "ttag";

export default function ProfileSummary(data) {
    return (
        <table className="styled-table">
            <tbody>
                <tr>
                    <td span={6}>
                        <strong>{t`Email`}</strong>
                    </td>
                    <td span={18}>{data.email}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Mobile`}</strong>
                    </td>
                    <td span={18}>{data.mobile}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`First name`}</strong>
                    </td>
                    <td span={18}>{data.first_name}</td>
                </tr>
                <tr>
                    <td span={6}>
                        <strong>{t`Last name`}</strong>
                    </td>
                    <td span={18}>{data.last_name}</td>
                </tr>
            </tbody>
        </table>
    );
}
ProfileSummary.displayName = "ProfileSummary";
