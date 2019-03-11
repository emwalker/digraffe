// @flow
import React, { Component, Fragment } from 'react'
import { Link } from 'found'
import classNames from 'classnames'

import type { UserType } from 'components/types'
import ViewerDropdown from './ViewerDropdown'
import SignIn from './SignIn'

type Props = {
  className?: ?string,
  viewer: UserType,
}

class Header extends Component<Props> {
  static defaultProps = {
    className: '',
  }

  get className(): string {
    return classNames('Header', this.props.className)
  }

  renderGuestUserNav = () => (
    <Fragment>
      <SignIn />
    </Fragment>
  )

  renderUserNav = (viewer: UserType) =>
    <ViewerDropdown viewer={viewer} />

  render = () => {
    const { viewer } = this.props

    return (
      <header
        className={this.className}
      >
        <nav className="flex-self-center">
          <h1 className="h3 text-normal">
            <Link
              to="/"
              className="text-gray-dark n-link no-underline"
            >
              <div className="mr-2 d-inline-block primary-logo">
                <svg
                  height="28px"
                  width="28px"
                  fill="#000000"
                  xmlns="http://www.w3.org/2000/svg"
                  xmlnsXlink="http://www.w3.org/1999/xlink"
                  viewBox="0 0 60 60"
                  version="1.1"
                  x="0px"
                  y="0px"
                >
                  <title>Digraph</title>
                  <desc>Never lose a bookmark again.</desc>
                  <g stroke="none" strokeWidth="1" fill="none" fillRule="evenodd">
                    <g transform="translate(-601.000000, -1070.000000)" fill="#000000">
                      <g transform="translate(601.000000, 1070.000000)">
                        <path d="M5,32 C5.552,32 6,31.553 6,31 L6,28 C6,27.447 5.552,27 5,27 C4.448,27 4,27.447 4,28 L4,31 C4,31.553 4.448,32 5,32" />
                        <path d="M20.2837,51.2041 L17.5697,49.9271 C17.0707,49.6941 16.4737,49.9061 16.2387,50.4061 C16.0037,50.9061 16.2187,51.5021 16.7177,51.7371 L19.4317,53.0151 C19.5697,53.0791 19.7147,53.1091 19.8577,53.1091 C20.2327,53.1091 20.5927,52.8971 20.7627,52.5351 C20.9977,52.0351 20.7837,51.4391 20.2837,51.2041" />
                        <path d="M56,25 C56,24.447 55.553,24 55,24 C54.447,24 54,24.447 54,25 L54,28 C54,28.553 54.447,29 55,29 C55.553,29 56,28.553 56,28 L56,25 Z" />
                        <path d="M55,31 C54.447,31 54,31.447 54,32 L54,35 C54,35.553 54.447,36 55,36 C55.553,36 56,35.553 56,35 L56,32 C56,31.447 55.553,31 55,31" />
                        <path d="M56,40.1016 L56,37.9996 C56,37.4476 55.553,36.9996 55,36.9996 C54.447,36.9996 54,37.4476 54,37.9996 L54,40.1016 C53.115,40.2816 52.326,40.7096 51.678,41.2906 C51.636,41.2506 51.605,41.2016 51.555,41.1676 L45.555,37.1676 C45.095,36.8616 44.474,36.9856 44.168,37.4456 C43.862,37.9056 43.986,38.5256 44.445,38.8316 L50.445,42.8316 C50.463,42.8436 50.482,42.8476 50.501,42.8586 C50.189,43.5096 50,44.2306 50,44.9996 C50,45.4446 50.077,45.8676 50.187,46.2776 L46.955,47.7976 C46.456,48.0336 46.241,48.6286 46.476,49.1286 C46.646,49.4916 47.006,49.7036 47.382,49.7036 C47.524,49.7036 47.669,49.6726 47.807,49.6086 L51.081,48.0676 C51.997,49.2346 53.404,49.9996 55,49.9996 C57.757,49.9996 60,47.7566 60,44.9996 C60,42.5856 58.279,40.5656 56,40.1016" />
                        <path d="M42.4307,49.9268 L39.7167,51.2038 C39.2167,51.4398 39.0017,52.0348 39.2377,52.5348 C39.4077,52.8978 39.7677,53.1098 40.1427,53.1098 C40.2847,53.1098 40.4307,53.0788 40.5677,53.0148 L43.2827,51.7378 C43.7817,51.5018 43.9967,50.9058 43.7617,50.4058 C43.5257,49.9058 42.9307,49.6938 42.4307,49.9268" />
                        <path d="M36.0972,52.9072 L34.7562,53.5382 C34.1262,51.4972 32.2452,50.0002 30.0002,50.0002 C27.7552,50.0002 25.8732,51.4972 25.2432,53.5382 L23.9032,52.9072 C23.4042,52.6742 22.8082,52.8862 22.5722,53.3862 C22.3372,53.8862 22.5512,54.4822 23.0512,54.7182 L25.0672,55.6672 C25.3972,58.1062 27.4722,60.0002 30.0002,60.0002 C32.5282,60.0002 34.6022,58.1062 34.9322,55.6672 L36.9492,54.7182 C37.4492,54.4822 37.6632,53.8862 37.4282,53.3862 C37.1922,52.8862 36.5962,52.6742 36.0972,52.9072" />
                        <path d="M39.5229,7.3555 L42.2969,8.4985 C42.4209,8.5485 42.5499,8.5735 42.6769,8.5735 C43.0699,8.5735 43.4429,8.3395 43.6029,7.9545 C43.8129,7.4435 43.5689,6.8595 43.0589,6.6485 L40.2849,5.5055 C39.7739,5.2985 39.1899,5.5405 38.9789,6.0495 C38.7689,6.5605 39.0119,7.1445 39.5229,7.3555" />
                        <path d="M17.3228,8.5732 C17.4498,8.5732 17.5788,8.5492 17.7028,8.4982 L20.4768,7.3552 C20.9878,7.1442 21.2308,6.5602 21.0208,6.0502 C20.8108,5.5392 20.2248,5.2982 19.7158,5.5062 L16.9418,6.6482 C16.4308,6.8592 16.1878,7.4432 16.3978,7.9542 C16.5568,8.3402 16.9298,8.5732 17.3228,8.5732" />
                        <path d="M23.7954,5.9082 C23.9224,5.9082 24.0514,5.8842 24.1754,5.8332 L25.0474,5.4732 C25.2914,8.0062 27.4054,10.0002 30.0004,10.0002 C32.5944,10.0002 34.7094,8.0062 34.9524,5.4732 L35.8244,5.8332 C35.9484,5.8842 36.0774,5.9082 36.2044,5.9082 C36.5974,5.9082 36.9704,5.6752 37.1294,5.2892 C37.3394,4.7782 37.0964,4.1942 36.5864,3.9832 L34.6464,3.1842 C33.9174,1.3252 32.1154,0.0002 30.0004,0.0002 C27.8844,0.0002 26.0834,1.3252 25.3534,3.1842 L23.4144,3.9832 C22.9034,4.1942 22.6604,4.7782 22.8704,5.2892 C23.0294,5.6752 23.4024,5.9082 23.7954,5.9082" />
                        <path d="M31,38.3818 L31,28.6778 L39,25.4778 L39,34.3818 L31,38.3818 Z M21,25.4778 L29,28.6778 L29,38.3818 L21,34.3818 L21,25.4778 Z M40.869,23.5108 C40.854,23.4858 40.835,23.4628 40.819,23.4388 C40.786,23.3908 40.751,23.3438 40.709,23.3018 C40.685,23.2778 40.658,23.2558 40.631,23.2338 C40.607,23.2138 40.587,23.1898 40.561,23.1718 C40.54,23.1568 40.516,23.1508 40.494,23.1378 C40.477,23.1278 40.464,23.1138 40.447,23.1058 L30.447,18.1058 C30.166,17.9648 29.834,17.9648 29.553,18.1058 L19.553,23.1058 C19.535,23.1138 19.522,23.1288 19.505,23.1388 C19.484,23.1508 19.46,23.1578 19.439,23.1718 C19.413,23.1898 19.393,23.2138 19.369,23.2338 C19.342,23.2558 19.315,23.2778 19.291,23.3018 C19.249,23.3438 19.214,23.3908 19.181,23.4388 C19.165,23.4628 19.146,23.4858 19.131,23.5108 C19.089,23.5848 19.057,23.6618 19.035,23.7438 C19.034,23.7478 19.032,23.7518 19.031,23.7548 C19.011,23.8338 19,23.9158 19,23.9998 L19,34.9998 C19,35.3788 19.214,35.7248 19.553,35.8948 L29.553,40.8948 C29.567,40.9018 29.584,40.8988 29.599,40.9048 C29.727,40.9618 29.862,40.9998 30,40.9998 C30.138,40.9998 30.273,40.9618 30.401,40.9048 C30.416,40.8988 30.433,40.9018 30.447,40.8948 L40.447,35.8948 C40.786,35.7248 41,35.3788 41,34.9998 L41,23.9998 C41,23.9158 40.989,23.8338 40.969,23.7548 C40.968,23.7518 40.966,23.7478 40.965,23.7438 C40.943,23.6618 40.911,23.5848 40.869,23.5108 L40.869,23.5108 Z" />
                        <path d="M9.5547,16.168 C9.5367,16.156 9.5157,16.155 9.4977,16.145 C9.8097,15.492 9.9997,14.771 9.9997,14 C9.9997,13.236 9.8147,12.521 9.5067,11.873 L13.0797,10.401 C13.5907,10.19 13.8337,9.606 13.6237,9.096 C13.4127,8.585 12.8277,8.343 12.3177,8.552 L8.6187,10.075 C8.4997,10.125 8.3987,10.197 8.3107,10.282 C7.4267,9.494 6.2747,9 4.9997,9 C2.2427,9 -0.0003,11.243 -0.0003,14 C-0.0003,16.414 1.7207,18.435 3.9997,18.898 L3.9997,22 C3.9997,22.553 4.4477,23 4.9997,23 C5.5527,23 5.9997,22.553 5.9997,22 L5.9997,18.898 C6.8857,18.719 7.6737,18.29 8.3227,17.708 C8.3647,17.749 8.3947,17.799 8.4457,17.832 L14.4457,21.832 C14.6157,21.945 14.8087,22 14.9987,22 C15.3227,22 15.6387,21.844 15.8317,21.555 C16.1377,21.095 16.0137,20.475 15.5547,20.168 L9.5547,16.168 Z" />
                        <path d="M45.001,22 C45.191,22 45.384,21.945 45.555,21.832 L51.555,17.832 C51.605,17.799 51.636,17.749 51.678,17.709 C52.326,18.29 53.115,18.719 54,18.898 L54,22 C54,22.553 54.447,23 55,23 C55.553,23 56,22.553 56,22 L56,18.898 C58.279,18.435 60,16.414 60,14 C60,11.243 57.757,9 55,9 C53.726,9 52.573,9.494 51.689,10.282 C51.602,10.197 51.5,10.125 51.381,10.075 L47.683,8.552 C47.172,8.345 46.588,8.585 46.377,9.096 C46.167,9.606 46.41,10.19 46.921,10.401 L50.493,11.873 C50.186,12.521 50,13.236 50,14 C50,14.771 50.189,15.492 50.502,16.145 C50.484,16.155 50.463,16.156 50.445,16.168 L44.445,20.168 C43.986,20.475 43.862,21.095 44.168,21.555 C44.361,21.844 44.678,22 45.001,22" />
                        <path d="M14.4453,37.168 L8.4453,41.168 C8.3953,41.201 8.3643,41.251 8.3223,41.292 C7.6743,40.71 6.8853,40.281 6.0003,40.102 L6.0003,35 C6.0003,34.447 5.5523,34 5.0003,34 C4.4473,34 4.0003,34.447 4.0003,35 L4.0003,40.102 C1.7203,40.565 0.0003,42.586 0.0003,45 C0.0003,47.757 2.2433,50 5.0003,50 C6.5953,50 8.0033,49.234 8.9193,48.067 L12.1933,49.608 C12.3313,49.673 12.4753,49.703 12.6183,49.703 C12.9933,49.703 13.3533,49.491 13.5243,49.129 C13.7593,48.629 13.5443,48.033 13.0453,47.798 L9.8143,46.277 C9.9233,45.867 10.0003,45.444 10.0003,45 C10.0003,44.23 9.8113,43.51 9.4993,42.858 C9.5173,42.848 9.5373,42.844 9.5543,42.832 L15.5543,38.832 C16.0143,38.525 16.1383,37.905 15.8323,37.445 C15.5253,36.985 14.9043,36.861 14.4453,37.168" />
                      </g>
                    </g>
                  </g>
                </svg>
              </div>

              Digraph
            </Link>
          </h1>
        </nav>
        <nav className="user-nav flex-self-center">
          <a
            className="text-gray-dark px-2"
            href="/wiki/topics/df63295e-ee02-11e8-9e36-17d56b662bc8"
          >
            Everything
          </a>
          <a className="text-gray-dark px-2" href="/about">About</a>
          {viewer.isGuest
            ? this.renderGuestUserNav()
            : this.renderUserNav(viewer)
          }
        </nav>
      </header>
    )
  }
}

export default Header
