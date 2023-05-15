import useNavigation from '../hooks/use-navigation';

interface Props {
  path: string;
  children: React.ReactNode;
}

function Route({ path, children }: Props): JSX.Element | null {
  const { currentPath } = useNavigation()!;

  if (path === currentPath) {
    // wrapped `children` in a fragment to ensure proper rendering
    return <>{children}</>;
  }

  return null;
}

export default Route
