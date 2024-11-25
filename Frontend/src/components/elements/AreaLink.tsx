import { Link, LinkProps } from '@mui/material';

interface AreaLinkProps extends LinkProps {
    text: string;
}

const AreaLink: React.FC<AreaLinkProps> = ({ text, ...props }) => {
    return (
        <Link
            {...props}
            sx={{
                ...props.sx,
                color: "black",
                textDecoration: "underline #000000"
            }}
		>
            {text}
        </Link>
    );
};
  
export { AreaLink };
