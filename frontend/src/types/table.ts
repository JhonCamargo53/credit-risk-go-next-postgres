export interface DataColumn {
    id: string;
    label: string;
    minWidth?: number;
    maxWidth?: number;
    className?: string;
}

export interface PageData {
    lastPage: number;
    limit: number;
    currentPage: number;
    total:number
}
export interface GenericTableRow<T> {
    index?: number
    data?: T
    pageData?: PageData;
}

export interface PaginationData<T> {
    data: T[],
    total: number,
    lastPage: number,
    currentPage: number,
    limit: number
}