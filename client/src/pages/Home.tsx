import { FC } from 'react'
import Navigation from '../components/Navigation'


const Home: FC = () => {
	return (
		<>
			<Navigation />

			<h1 style={{ fontSize: '2.5rem' }} className="text-center">Pacstall - The AUR for Ubuntu</h1>
			<hr className="uk-divider-icon" />


			<div className="uk-card uk-card-default uk-card-body uk-width-1-2@m uk-flex uk-flex-center uk-flex-wrap uk-card-hover uk-align-center"
				style={{ textAlign: 'center' }}>
				<h3 className="uk-card-title">Why is this any different than any other package manager?</h3>
				<p>Pacstall uses the stable base of Ubuntu but allows you to use bleeding edge software with little to no
					compromises, so you don't have to worry about security patches or new features.</p>
			</div>
			<div className="uk-card uk-card-default uk-card-body uk-width-1-2@m uk-flex uk-flex-center uk-flex-wrap uk-card-hover uk-align-center"
				style={{ textAlign: 'center' }}>
				<h3 className="uk-card-title">How does it work then?</h3>
				<p>Pacstall takes in files known as <a
					href="https://github.com/pacstall/pacstall/wiki/Pacscript-101"></a> (similar to PKGBUILD's) that contain the necessary contents to build packages, and builds them into
					executables on your system.</p>
			</div>

			<div>
				<img
					className="uk-border-circle uk-align-center"
					src="/public/pacstall.svg"
					width="200px"
					height="200px"
					alt="Pacstall logo"
					loading="lazy" />
			</div>
		</>
	)
}

export default Home
