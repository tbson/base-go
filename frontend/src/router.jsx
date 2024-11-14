import { createBrowserRouter } from 'react-router-dom';
import NotMatch from 'component/common/route/not_match';
import PrivateRoute from 'component/common/route/private_route.jsx';
import BlankLayout from 'component/common/layout/blank';
import MainLayout from 'component/common/layout/main';
import AccountWrapper from 'component/account/index';
import AuthWrapper from 'component/auth/index';
import ConfigWrapper from 'component/config/index';

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
                        element: <AuthWrapper />,
                        children: [
                            {
                                path: '',
                                lazy: async () => ({
                                    Component: (
                                        await import('component/auth/login')
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
                                element: <AccountWrapper />,
                                children: [
                                    {
                                        path: '',
                                        lazy: async () => ({
                                            Component: (
                                                await import(
                                                    'component/account/profile'
                                                )
                                            ).default
                                        })
                                    },
                                    /*
                                    {
                                        path: 'account/user',
                                        lazy: async () => ({
                                            Component: (
                                                await import('component/account/user')
                                            ).default
                                        })
                                    },
                                    {
                                        path: 'account/role',
                                        lazy: async () => ({
                                            Component: (
                                                await import('component/account/role')
                                            ).default
                                        })
                                    }
                                    */
                                ]
                            },
                            {
                                path: '',
                                element: <ConfigWrapper />,
                                children: [
                                    {
                                        path: 'config/variable',
                                        lazy: async () => ({
                                            Component: (
                                                await import(
                                                    'component/config/variable'
                                                )
                                            ).default
                                        })
                                    }
                                ]
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
