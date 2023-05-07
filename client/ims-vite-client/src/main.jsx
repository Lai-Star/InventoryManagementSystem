import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import { NavigationProvider } from './context/navigation.jsx'

const el = document.getElementById('root')
const root = ReactDOM.createRoot(el)

root.render(
    <NavigationProvider>
        <App />
    </NavigationProvider>
)
