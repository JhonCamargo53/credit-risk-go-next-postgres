import { create } from "zustand";
import { PageData } from "@/types/table";
import { CreditStatus } from "@/types/creditStatus";
import { fetchCreditStatuses } from "@/services/creditStatusService";

interface CreditStatusStore {
    creditStatuses: CreditStatus[];
    allCreditStatuses: CreditStatus[];
    selectedCreditStatus: CreditStatus | null;

    pageData: PageData;

    loading: {
        fetching: boolean;
    };

    error: string | null;

    resetStatus: () => void;

    handlePagination: (pageData: PageData) => void;

    fetchCreditStatuses: () => Promise<void>;
}

export const useCreditStatusStore = create<CreditStatusStore>((set, get) => ({
    creditStatuses: [],
    allCreditStatuses: [],
    selectedCreditStatus: null,

    pageData: {
        currentPage: 1,
        limit: 10,
        total: 0,
        lastPage: 1,
    },

    loading: {
        fetching: false,
    },

    error: null,

    resetStatus: () => set({ error: null }),

    handlePagination: (pageData: PageData) => {
        const all = get().allCreditStatuses;
        const { currentPage, limit } = pageData;

        const start = (currentPage - 1) * limit;
        const end = start + limit;

        const paginated = all.slice(start, end);

        set({
            creditStatuses: paginated,
            pageData: {
                currentPage,
                limit,
                total: all.length,
                lastPage: Math.ceil(all.length / limit),
            },
        });
    },

    fetchCreditStatuses: async () => {
        try {
            set((state) => ({
                loading: { ...state.loading, fetching: true },
                error: null,
            }));

            const data = await fetchCreditStatuses()

            set({ allCreditStatuses: data });

            get().handlePagination(get().pageData);
        } catch (err: any) {
            set({ error: err.message || "Error al obtener estados de crédito" });
            throw err;
        } finally {
            set((state) => ({
                loading: { ...state.loading, fetching: false },
            }));
        }
    },
}));
