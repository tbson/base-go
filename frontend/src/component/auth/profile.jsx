import * as React from "react";
import StorageUtil from "service/helper/storage_util";
import AdminProfile from "component/admin/profile";
export default function Profile() {
    const userType = StorageUtil.getProfileType();
    if (userType === "admin") {
        return <AdminProfile />;
    }
    return null;
}

Profile.displayName = "Profile";
