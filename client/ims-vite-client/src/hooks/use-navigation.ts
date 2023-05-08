import { useContext } from 'react';
import NavigationContext, {
  NavigationContextType,
} from '../context/navigation';

function useNavigation() {
  return useContext<NavigationContextType | null>(NavigationContext);
}

export default useNavigation;
