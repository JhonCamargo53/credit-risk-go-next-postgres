/**
 *  Listado de roles
 */
export const ADMIN = "ADMIN";
export const USER = "USER"


export type AvailableRoles =
    typeof ADMIN |
    typeof USER

/**
*  Diccionario de Roles
*/

export const roleDictionary: { [role in AvailableRoles]: number } = {
    [ADMIN]: 1,
    [USER]: 2,
};