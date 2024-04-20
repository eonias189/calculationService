import { FC } from "react";
import "./loader.css";

interface LoaderProps {
    className?: string;
}

const Loader: FC<LoaderProps> = ({ className }) => {
    return (
        <span
            className={`border-4 border-teal-500 border-b-transparent block size-32 rounded-full rotate ${
                className ?? ""
            }`}
        ></span>
    );
};

export default Loader;
