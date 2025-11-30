import { AvailableRoles } from "@/config/role.config";
import { routeAccessConfig } from "@/config/route.access.config";
import { IRoute } from "@/types/route";

export const getAllowedRolesByPath = (path: string): AvailableRoles[] => {

    const segments = path.split("/").filter(Boolean);
    
    const findRoles = (segs: string[], routes: IRoute[]): AvailableRoles[] => {
        if (!segs.length) return [];

        const [currentSegment, ...rest] = segs;

        const match = routes.find(
            r => r.path === currentSegment || (r.path?.startsWith("[") && r.path?.endsWith("]"))
        );

        if (!match) return [];

        if (rest.length === 0) {
            return match.allowedRoles || [];
        }

        return findRoles(rest, match.subpaths || []);
    };

    return findRoles(segments, routeAccessConfig);
};
