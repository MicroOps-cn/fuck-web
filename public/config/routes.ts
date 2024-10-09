import { loginPath, forgotPasswordPath } from './env';

const routes: Route[] = [
  {
    path: '/account',
    layout: false,
    routes: [
      {
        path: loginPath,
        component: './Account/Login',
        access: 'canAnonymous',
      },
      {
        path: forgotPasswordPath,
        component: './Account/ForgotPassword',
        access: 'canAnonymous',
      },
      {
        path: '/account/resetPassword',
        component: './Account/ResetPassword',
        access: 'canAnonymous',
      },
      {
        path: '/account/activateAccount',
        component: './Account/ActivateAccount',
        access: 'canAnonymous',
      },
      {
        path: '/account/settings',
        component: './Account/Setting',
      },
      {
        path: '/account/events',
        component: './Account/Events',
      },
      {
        component: './404',
      },
    ],
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    icon: 'dashboard',
    access: 'canViewer',
    component: './Dashboard',
  },
  {
    name: 'roles',
    icon: 'table',
    path: '/roles',
    access: 'canViewer',
    component: './Roles',
  },
  {
    name: 'users',
    icon: 'UserOutlined',
    path: '/users',
    access: 'canViewer',
    component: './User',
  },
  {
    name: 'events',
    path: '/events',
    component: './Events',
    icon: 'SecurityScanOutlined',
    access: 'canAdmin',
  },
  {
    name: 'settings',
    path: '/settings',
    icon: 'setting',
    access: 'canAdmin',
    component: './Setting',
  },
  {
    path: '/403',
    layout: false,
    access: 'canAnonymous',
    component: './403',
  },
  {
    path: '/warning',
    layout: false,
    access: 'canAnonymous',
    component: './Warning',
  },
  {
    path: '/',
    redirect: '/dashboard',
  },
  {
    component: './404',
  },
];

export interface Route {
  path?: string;
  component?: string | (() => any);
  wrappers?: string[];
  redirect?: string;
  exact?: boolean;
  routes?: Route[];
  access?: 'canViewer' | 'canUser' | 'canAdmin' | 'canEditor' | 'canAnonymous' | 'forbidden';
  [k: string]: any;
}

export default routes;
