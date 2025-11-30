import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

dayjs.extend(utc);
dayjs.extend(timezone);

export function utcToLocalTime(date: string | Date, format: string = 'DD/MM/YYYY', timezone = 'America/Bogota') {

    if (!timezone) {
        timezone = 'America/Bogota'
    }

    return dayjs.utc(date).tz(timezone).format(format);
}

export const getCurrentLocalTime = (timezone = 'America/Bogota') => {
    return dayjs().tz(timezone);
}

export function formatDateToDDMMYYYY(date: string | Date): string {
    if (!date) return '';
    return dayjs(date).format('DD/MM/YYYY');
}