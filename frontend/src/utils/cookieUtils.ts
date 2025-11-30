export function getCookieValueService(cookieName: string) {
    const cookieRegex = new RegExp(`(?:(?:^|.*;\\s*)${cookieName}\\s*=\\s*([^;]*).*$)|^.*$`);
    const cookieMatch = document.cookie.match(cookieRegex);
    if (cookieMatch && cookieMatch[1]) {
        return cookieMatch[1];
    }

    return null;
}

export const serviceSetCookie = (name: string, value: any, secondsDuration: number) => {

    let d = new Date();
    d.setTime(d.getTime() + (secondsDuration * 1000));
    let expires = "expires=" + d.toUTCString();
    document.cookie = name + "=" + value + "; " + expires;
}

export function getCookieExpirationTime(cookieName: string) {
    let name = cookieName + "=";
    let decodedCookie = decodeURIComponent(document.cookie);
    let cookieArray = decodedCookie.split(';');

    for (let i = 0; i < cookieArray.length; i++) {
        let cookie = cookieArray[i].trim();
        if (cookie.indexOf(name) == 0) {
            let cookieValue = cookie.substring(name.length);
            alert(cookieValue)
            return cookieValue;
        }
    }

    return null;

}