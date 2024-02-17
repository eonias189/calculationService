"use client";
import { useState } from "react";

export type Fetcher<T> = {
    isLoading: boolean;
    data: T | undefined;
    update(): void;
};

export const useFetcher = <T>(fetch: () => Promise<T>): Fetcher<T> => {
    const [data, setData] = useState<T>();
    const [isLoading, setIsLoading] = useState(false);
    return {
        isLoading,
        data,
        update: () => {
            setIsLoading(true);
            fetch()
                .then(setData)
                .then(() => setIsLoading(false));
        },
    };
};
