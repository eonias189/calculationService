"use client";
import { useEffect, useState } from "react";

export type Fetcher<T> = {
    isLoading: boolean;
    data: T | undefined;
    update(): void;
};

export const useFetcher = <T>(fetch: () => Promise<T>): Fetcher<T> => {
    const [data, setData] = useState<T>();
    const [isLoading, setIsLoading] = useState(false);
    const fetcher = {
        isLoading,
        data,
        update: () => {
            setIsLoading(true);
            fetch()
                .then(setData)
                .then(() => {
                    setTimeout(() => setIsLoading(false), 500);
                });
        },
    };
    useEffect(() => {
        fetcher.update();
    }, []);
    return fetcher;
};
