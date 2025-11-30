import { create } from "zustand";
import { PageData } from "@/types/table";
import { Asset } from "@/types/asset";
import { fetchCustomerAssets } from "@/services/customerAssetService";
import { fetchAssets } from "@/services/assetService";

interface AssetStore {
    assets: Asset[];
    allAssets: Asset[];
    selectedAsset: Asset | null;

    pageData: PageData;

    loading: {
        fetching: boolean;
    };

    error: string | null;

    resetStatus: () => void;

    handlePagination: (pageData: PageData) => void;

    fetchAssets: () => Promise<void>;
}

export const useAssetStore = create<AssetStore>((set, get) => ({
    assets: [],
    allAssets: [],
    selectedAsset: null,

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
        const all = get().allAssets;
        const { currentPage, limit } = pageData;

        const start = (currentPage - 1) * limit;
        const end = start + limit;

        const paginated = all.slice(start, end);

        set({
            assets: paginated,
            pageData: {
                currentPage,
                limit,
                total: all.length,
                lastPage: Math.ceil(all.length / limit),
            },
        });
    },

    fetchAssets: async () => {
        try {
            set((state) => ({
                loading: { ...state.loading, fetching: true },
                error: null,
            }));

            const data = await fetchAssets();

            set({ allAssets: data });

            get().handlePagination(get().pageData);
        } catch (err: any) {
            set({ error: err.message || "Error al obtener assets" });
            throw err;
        } finally {
            set((state) => ({
                loading: { ...state.loading, fetching: false },
            }));
        }
    },
}));
