import { createBrowserRouter } from 'react-router-dom';
import StorageUtil from 'service/helper/storage_util';
import NotMatch from 'component/common/route/not_match';
import PrivateRoute from 'component/common/route/private_route.jsx';
import BlankLayout from 'component/common/layout/blank';
import AdminLayout from 'component/common/layout/admin';
import UserLayout from 'component/common/layout/user';
import TenantLayout from 'component/common/layout/tenant';

const userInfor = StorageUtil.getUserInfo();

const MainLayout = userInfor?.profile_type === 'admin' ? AdminLayout : UserLayout;

import App from 'src/app';

const router = createBrowserRouter([
    {
        path: '/',
        element: <App />,
        children: [
            {
                path: 'login',
                element: <BlankLayout />,
                children: [
                    {
                        path: '',
                        lazy: async () => ({
                            Component: (await import('component/auth/login')).default
                        })
                    }
                ]
            },
            {
                path: 'account/tenant/:tenant_id',
                element: <TenantLayout />,
                children: [
                    {
                        path: '',
                        element: <PrivateRoute />,
                        children: [
                            {
                                path: '',
                                lazy: async () => ({
                                    Component: (
                                        await import('component/account/tenant/detail')
                                    ).default
                                })
                            },
                            {
                                path: 'role',
                                lazy: async () => ({
                                    Component: (
                                        await import('component/account/role')
                                    ).default
                                })
                            }
                        ]
                    }
                ]
            },
            {
                path: '',
                element: <MainLayout />,
                children: [
                    {
                        path: '',
                        element: <PrivateRoute />,
                        children: [
                            {
                                path: '',
                                lazy: async () => ({
                                    Component: (
                                        await import('component/account/profile')
                                    ).default
                                })
                            },
                            {
                                path: 'config/variable',
                                lazy: async () => ({
                                    Component: (
                                        await import('component/config/variable')
                                    ).default
                                })
                            },
                            {
                                path: 'account/auth-client',
                                lazy: async () => ({
                                    Component: (
                                        await import('component/account/auth_client')
                                    ).default
                                })
                            },
                            {
                                path: 'account/tenant',
                                lazy: async () => ({
                                    Component: (
                                        await import('component/account/tenant')
                                    ).default
                                })
                            }
                        ]
                    }
                ]
            },
            {
                path: '*',
                element: <NotMatch />
            }
        ]
    }
]);
export default router;
