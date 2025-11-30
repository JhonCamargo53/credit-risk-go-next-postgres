import {
  IoHomeOutline, IoPeopleOutline
} from "react-icons/io5";
import { FaAddressCard} from "react-icons/fa";
import { IRoute } from "@/types/route";

export const routeAccessConfig: IRoute[] = [
  {
    path: 'home',
    label: 'Home',
    icon: <IoHomeOutline />,
    allowedRoles: ['ADMIN', 'USER'],
    renderNavigator: true,
  },
  {
    path: 'users',
    label: 'Usuarios',
    icon: <IoPeopleOutline />,
    allowedRoles: ['ADMIN'],
    renderNavigator: true
  },
  {
    path: 'customers',
    label: 'Clientes',
    icon: <FaAddressCard />,
    allowedRoles: ['ADMIN', 'USER'],
    renderNavigator: true
  },
];
