package server

import "text/template"

var indexTemplate *template.Template = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html lang="ja">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>ikusan</title>

		<!-- Bootstrap -->
		<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">

		<!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
		<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
		<!--[if lt IE 9]>
			<script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
			<script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
		<![endif]-->
		<style>
			header {
				text-align: center;
			}
			.affix {
				position: fixed;
				top: 50px;
				width: 21%;
			}
		</style>
	</head>
	<body data-target="#sidebar" data-spy="scroll">
		<div class="container">
			<header>
				<h1>ikusan</h1>
				<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAmYAAAEHCAYAAAAEbIroAAAZ70lEQVR4nO3d7XKkLLsGUHlrn/8pu39Mdx5j/EBFucG1qlJjJnZLo+LViJrGcRwAAKjvf7ULAADAP3+CWUop+8XmNe+b5211WUfea+99UkpjSmmx2/37t1LLO+OudVTidUvvE70+gfv9X+0CALFMDvzjMAzfX8ZxHOeJYPr3xbfKnK9r6hM4QjADlvwKEcNCEFgIFvO/fycX55v0DGWFjMbHw6pPIIsxZsDPabTPr7shotRiPz/j56cb6hM4S48Z8Mus5+aJU2bT0PLUMh+jPoEjBDNgGGKMW+opUKhP4BTBDNgd3/Sw5gOF+gTOEsyAqASKstQnNEAwA6ISIMpSn9AAwQyIRoAoS31CQwQzIAoBoiz1CQ0SzIDaBIiy1Cc0TDADajkVID43bk3D4O71M+oTOiCYwQt8nte49HzGai6U5Xt3+2qfRX0Cd/FIJqBF3zBBGeoTgtBjBu/QTY/I3sO8nypG5eUXE6Q+gQ/BDCZSSovjbGqPvVlb/lJ5V8p/6qD7OWV3xjiOY1qrz7VllZ53HMdD8+Yu9w31OR17dlTt/QVa5lQmUFo3vUlB1KpPpzehAj1m0JmNXpknBqs3Hcoq193icoe69enCAHiYHjN4hycOrr0ewGt9rij1qecMHiSYASVECRG9iFafwhk8RDCD/t19kI8WIkqq8dmi1qdwBg8QzIArooaIVkWvT+EMbiaYAWdFDxGtaaU+hTO4kWAGfbvrYN9KiLjiyc/YWn0KZ3ATwQw4qrUQEV2r9SmcwQ0EM+jXHQf8VkPEUU99ztbrUziDwgQzIFfrISKaXupTOIOCBDPoU+mDfi8hIoeb8R4nnEEhghmwp7cQUVuv9SmcQQGCGfSn5IG/1xCxxs14rxHO4CIPMYfOlHzYdqUHd1dz9+dtqT7H8XS+SsMwDCmlVwe0ltY1saQLOx8AAAX9OZWZUn7IN695e5u3h2WVfO3R11/9nDWULF/U7frIvAVeP6aUhjM/vfvUzfiGz8p5xpgBUJJxZuvSoH7YYYwZvNzGt/exh3EyJT9f73VV0Dd8qJOJydChxXqZjMvLqjdDkfqkxww6U2jQde8HVVeu3k/P0HHTHjV191KCGQB3Ec7OEdBeTDAD5nrvAdJb9izh7DwB7YUEMwDuJpxdI6C9iGAGwBOEs+sEtBdwVSYw1fupOacx63K15jWHrtqkTYIZAE8Szo4TyF5EMAO+ej9Y6i2LQzjLI5C9kGAG/dGI0wLhbJ1A9mKCGTAM/R8g9ZbFJJz9JpAhmEGHsg50KSWPESIC4ezD/sgwuF0GdCe3cZ/M1/tBUW9ZfG6lAR96zODlIn1LL/Sczx/jOKaSny9SXUV14cHaaRjKbwO12WY4Knk6PQBADH96zFJK2d94zGveFuZt6fWl1B4/llIahs+pqRrluKvOP5/rjHEcx3S1N+j7HnphoF9OZQLZJsFkOtZqKSgYi/Wfn7oQqIA9ghl0bqOX52zPy69QNiwEsLsDyJnPdEM95HgkoE564nIv/LixNMAVghkEMgsPiwfbQgfV7MDwPS35CS+7oayiM+W58zM8WT/TdTL9HWiMYAYxfQ/qobo2Zj1LDv7raoVWAQ0a5z5mENsd4exoaIjWM7YkUm9ZhPpKw3/bTqhwD2zTYwbxPHpgX7tS8Ns7ZsD6IRFC2ZQeNGiMHjOIZenAfvcpzbTwc9/CJkHwynStsm2IFsqm9KBBI/SYQSx3H9i3wsN0XNud5Vgbp3Z0+qvkacwjy71ahifpMYNGCGbQhqce9Bw9YEQUuc4EMmiMYAbvkRMgWjuA1x70HzWUCWTQKMEM2vGr1+wz7uk7QL9isV4rYigTyKBxghkEkhGw5uOecsPB3nwRQ8aemr1l0eqraCAT+qEewQza9tTYs5LWnh5wdDpC2UK4ckuT6UPRJ9MtblfQBcEMGjXpybh6xeCjB+BpiLgwfUtv2U3LDW3tMw/CGVQhmEHn9npTWryB7Jkyl/icLdbVGQdCP1BYMn4AACCGP3f+Tyn/C5J5zdvCvLVfX/I1K+UbU0rDmZ/WHa27kz/jmWXd+BnGu5+AANTjkUzQvrsf2fRmUcdYRSwTUIAxZnCzjR6RseCYpS4Haj9Ud6vLGDqrTyA+PWZQxx0H/bf0nHk0VWXTU7xAWYIZ9OUt4exOQhlQjWAGz7v7wN9zOHvyBrOsU09wE2PMoE9djjm7mfrK9Jb7uUENesygX9k9Z9PxQtGmzzjxXq2Esj/ljLBejDeDcgQzeNbTASA3nM0fxRNp+utI3e2919n3rW1pfdZeL1v/DxwkmEH/eh5zdlVLoQx4AcEMnlMzBLQezu6oO6EMCEcwg/doPZyVJJQBIQlm8IwoQaDFcFa67qKsC4A/BDN4n6VwNgaePmPrfXsLZRHWy8+D3l2hCdcIZnC/iGHgVzib3pcq2PSpuiv5XkFVX3e/CuO+ZlCMG8zCze4+aI3j6Q6K8AfTknXXQ3j4rOvIn6On8AtVpAuNOgAABf05lZlS/pcd8/Yzb9Ry9TLv08ugnDPr4M59LKU0fn4OlwuIzxgzgLakoc2ra4EMxpgBiyZX12V1zdw5LOJA79D4KUu33UmTeu72M8KbCWbwkJTS2FhgmF7NOP09ohbKeKu121Q0ts3B6wlmwJ7IAS1imWpZe/B7GJ/w+KucLkCD3wQzeE7r4SFSQItQBo77jo2z3mCFYAadWOqNmCvUO1EzoAlk+aIGIOEMNghm8Jy7D0Y5V+qVXP5qQJuFxM0yTcdATV43H8R/WyA7Mzar1fFc03GOT04vFWUQzmCRYAadyL0r/NazDA8Gi72w9OvvB3rrlg7YpwPa3uc9E6aiB7ANqdL0WlmEM5gRzOBlCoSK3HB0djlbrzsc0BoOUVcIPNAowQzIFWl8V6QLEci00GtqvcGMYAbsiRx+BDSgK4IZsKalsCOg/cdpTGiYYAYsijQ268CFA2HKHNT8atmnprN9Hr/V2lMyoBjBDB7ypgPN2c/6pjoqLecZmtP6fXIayPe/2gUAgAmnYnm15DllAAAx/DmVmVLKHs9h3tjzRi1XiXlLqbFM9n3HGQ3DnycD3LXtNb0RTO+2//mvX1+6J/WZxnHc/b03T+3n2hNKMMYM+PE5QA/DbBD35MB/698nHj2d1dF4qK2B9/M77e/9DlQgmAFze1fW3f33noLSozYuAPjW87ze134HKjH4Hw5KKVU79TVd9h3TQ4BQxi2+vWFZv5/YboBC9JgBP2Y9VXs9WcX/zj3mD7jf+x2oR48ZHFezpyfMjUDpmu0GKtFjBgfVHP/kRqA8odR28znVuTmfqxjhN8EMLrg6xmYcx5RSml61uDsNEa1sr/OxbCsvXd6PFm4Dsvh36IlgBjNbYWt+ICh0YFgbd2U8Fi35s73mjl3b24/2/n5kn4XoBDOY0ZBDW+yz9MTgfwCAIAQzAIAgBDMAgCAEM5qzdffxRu9EPh6choia217X2o6laU9A4CmCGV1pcRCwe43Rgxa3V/seEbkqk9Dm30Q/jeTu3ceXXrfyXgA5ij4BQXvEmuSuywAAMfzpMUspZT8iw7z9z3uXlP7df3IYfFN8wh3rvNZ2tLbt3Li/NP3tdeHu+b++kJ9dj621Ob1YW4/0w6lMivocNIdh1tW/EL48AJlfom47HX1x8PDxPliPnRPMuMPu+IuODnaUZdu5yaR3Rf01zHrsn6syOWXjcvHLg2LpT+atBmw7wbk1BNxPjxlFzXozHFjJZtsB0GPGeXo3OKLorQaoxvqCm+kx4xQ3YOQIN+/sg/UF99Njxo+U0vgZQ/KdHrf+/85ymN6+TUPuY6murKvvup7/nH0/KCHC/hilHZjr4PF0DHrM+OvnHjmZ/3+HtbFGb55ecuk9thrucRyT3hGCirA/RmoHpozT7IBgxtzazmwn74zgBRCPU5kAAEHoMQMglM9p9n/P3vLYIV5GMAMgmjS4NcethN+4nMokotH0n+klJd4DovqGs7kI+2OkdmBq8XUrV2uu1S+V/Xk6fUopOz2bt/95ienzwO9xHMc0nd6av/Q6j7Yd2QfOOVsX6rttn3ZjGHbaDp7nVCYEN7+txacRnZ7mccoHOMTD0OP602MGAEAdf3rMIpw+izpv1HKVmLfE6yhbd5NTDUf9nOK8c1uJsN1GmPcz/59ey73XRyh7g/Oe7h3Wpu2b1686q8OpTKhkOsZjmJyWvDjew2lNeuZqTbonmEFdJceKFT1gbfTYGSxMTcLZfdRtAG6XwTAMw58HYVPOxgPHw4ayAMuBLW71QLf0mEEls16nFkIZRKJ35x7qtTI9ZnzZEe9z560tnlxvthGi0XNGd/SYMQzDn94bCprWbeF6FpRimd5Rnefo4SlPnVakx4xbrDwCpJvp+Zi8CuPzNJrxpEEPTi2X6712mxJlmvoEM+6yNn7qDdN3qxHKBMF8wlkdV+s9QtsRYZrKBDNoi4DUBuuoDqG4DO1MRYIZtKNWY6mRpiXCGU0z+B/aIBw1avqYG4+4eYzB6+ept8r0mJElpTR+BrzXLsob1WwoNdLX6cGpQ73TJMGMXEcbufHF0yUJRn0QEur4U+8bVyOa1t6E4FQmR2SfHli7d1dP02uft6DajWTt5XdhcvpSXdbxq92K0HZEnqY+wYwsDi5/3B5aajeWtZffEnX1jAtj9KyfDLbjGJLBqAAAMfzpMUspZX8raW3enPcaPj0h4zguDXT/00tSI9hGqMu9eUuul9IuXMAwjuOY7t4+r75/ifKllMbvt+cWtreceadXR37tvefR9Vey3pbKm+vbfkVdF1HnbdXVNq1kWbjudacy1xq7z477/ZuxNczZJniatuhmB8JvjwHGthXU64JZbXsNQe/f7BpVtAHb+HbbY+PPNcIZd7BNBeZ2Gb/lNILj5OfKMtZnmD1QtpUHzM4f7N2JpxowDSVr3GrjPjl129u+2dvn6Y4es4lPb9V8g53vtCU2aN+C22AdEYU2gxJsQw3QY7bf+5UmPyUtvV/r34p72umf/Cw91Rv30XN2j6167Wnf7OmzdO2tPWbTnbDIhpoziLT38WMdjY/SgPXruxO2un71nHGGbaYhb+0xu9ILNs5+pu85/9vavGvv++odZ+1RKXdMb3j9eujcd79v+VtS6+WPaKlON9uCJ9srbdq7vLXHLNfaRr92uw0b/zXpweklNRowjWYdrfc8tV7+HjzZXrXUpnGRYLbtqXFgdp76rIP3aX19C2dssW006q2nMiOx89RXax1Y91zltGY507psfd9svfyvpsesrrWdZ/5/3exg04skglwMoQEj4nZ5hJ4zpmwLjdNjVs9bd55I3/BrroO3rv+oIm2XZ7Re/ihaD7ktl52PNwaztasnn/Jd3t7TBbZ+b93aQWR8eFoDxlTr4ab18rfm6fYqZ1qb1oHXncqsffVkzr2+5vN0dH+w6WmipStb00PTtRuw2stnZmu7bEzrPT6PWjltnVV3D7ZXLbRpFPS6YHbElZ2WuGoH3drLf6vW6v3CWLemPifXtbZtsy01ONAVAKBLbxxjRgApPf8Fb2mZR8pRYt7ZXbtvWfYT7/V0vV2dt+Rre19WhPV1R53VaHPgDKcyV6zsxIvn8fU6AgAlCGaEM+lVyh2E+33dxizvGYOhHgDa5VRmvm9vme6x+6Xh921NrnC10j+vrofMhz03tyygP3rMiGx6Kfj0d240vQv+l9P1/9xRN0/Wt3UL8ekxyzNtzPSaPe9sD9qre4kmjtaDbXzdHXXzZH1btxCcYEZLSp7iZJsD+DrhDLiNU5m0yCnOZ7iL/Lo76mbzPSenIeeh6szD161bCEqP2b6lxss3zrqmzxtdO7A46PxztR5s6+tq9JxNt/uUMf+VZQEV6DGjJXrIeLulbV/vF3REMNu21dhpDJ9zNJBZL/+cqoeFU2Lq8uOOunmyvq1biE8wu9n08nSXpR+mh4xSntyGbK/AacaYrcvpbcgZo2Ecx3E5Y8i2XuvAqB4AmqTHrLDP43Dmj75x2vMAjw3iBk/uf/Z14DTBbMWBcLA53+T0pYb6fg6I/6iHmSfD/l3L2nlf6xs68SeYvf05bwUa1ayD4tvrebjhQKKn7R/1QI8KtJl/9ovS7fA4jknb3pca7WkyIB0AIAY9ZjOl0/Hb6/Oqs+sjpZR9FezevEt/P/L+V5Zd670ylnX6dOk4jkXXT868KaWfcZ9PL/vqvJ/55+NWd+e7sj1c/Sw31pv2lGqe6j37E8ycBvlt0hBs3vJibT71Sadc0MLjtKe8gcH/+6YPzd4bfJszHwV8rn5dktWzQBHCGUBhgtnMxu0uhmFyEMqdj0eFr/fcU1Jrr136/8pBVDjrwJXtEihLMNsxO3W51U2TNR9csLRdZY+5WRoXVmhMmnAGUIg7//+Ve4BxIIrF+th351MoPOECoAA9ZjO53fm6/QngTBi9s3dLzxnARYIZvM9mgJqc8pz3gG1emZzz3pQxuy2G8WHQEacy6YEgcNzeqcfpQ+RTxvxH3pvrli46AjogmEGb7g6jS+8tnAHczKlMWqe3LNPCKci7681pTYCD9JjBs94WUvScteFt2yWEJZjRsrf2xrT2uYUzgEyCGeFN73j/1EOMU0pFlzV5r7cGFOGsrHFlusT7ARUZY0Z401sBTKab6jX6lvvKbQ0mY8Ru/dw7ZbyybGPOClnZJ4q8H1CXYEaTHjiQCBAbLjzKSZ0CbEiFnpUHAMBFf8aYpZT/hda85o0y75H3y1zmdFzbkdcdXc6h+Y++5kR5sr6p7dXPwTKOnzF9h39yl/XE+mxhvWS83rwF5n1imy5V1hLz3imz3saV6drFP8WpTKA2Y8/ojW26nl91P3nE3JLp4+fCnEEUzHidjW9RIZ45GL18N3Ege8BLt61autumJ9vPXoL5CTkr29zS63/NOAtJOYlpPk+zdS+Y0aQbHtz8yE58odzNNjIHdHcga8Qtdf45IL898PW6TZ/9PN/wtFQv078dXebZOk6zf0MQzIBIej2Q8V626X/26iDnVki7PW2ZywpNMIP4O3H08pXmQPacO+vZOvzP27fpUp99qzetm/oVzICIumpo3+jlpzCX2KbzbNVTTo/Z5TqeXN28d6p1HIa/2/rV13skE7CqxuOwposfPCpoUeX1wnlv3KaXgtKV8JQmP/P/n75vs/Wsx4y3i/4Ntnb50sr0k8uvXQcRlVgv6rWO1rfp2mVfC1y/bn1x8n2/V6wsvT73/y6/XjADomv9QAZzrW7TNco8Xebe/ci+8zTbWzYMTmXybtEbxujle1LzjW0wtq363rpN7217OX+fTy+d1jyyfYfaHwQzWhVmJzqo1XJH8NYDGf16wzZ99TPu3WLjav2FCmXDIJjxXuF2xpno5avlDQeyu9m2Ymlpmz5b1u/r9u5RduRqzOl7z+cZZ39fev1eeaoRzGhVKw3ZMAy/rpxrqtzDemNXU9bBofOrFSOul5+HSNcuR4NaDGd7P0uvWwtd48bfr5Rxq8wll1eUwf806eI9kh7/lvQtb2a5w3yLm5Y32H2pdgdPBytvURfWS5htiz/CXxAweX5ldhkzHgy+917zOjkUYs+UuTbBjNeJfsCOXr7SMhruNa+qpxIe2LZCB4un2KaLyr7NxIl5QkoXNiAAAAr6M8YspfyQaV7zRpv3rvdcmv/Ee0zv1n7otWeXe6WMufNFWO89z1vK3eWLUFdPzvvwssbP+L3DP7THqUya9gkIWeO23naKkH7Mn73nTMfrhB9/RjmCGUB8u7cb2OgdKfKl5O73Z5dw9hKCGa3LaawON2YppccPNisHvl9l11PyatP7NaVh+NlmtrbVuw/kgsKzhLMXEMwgnqXnvmmIX2wWyG0L76ZN6JwbzNKDrfvaaMDo2d6d1PWW9amlG9JykB4ziM83ZBYZ2/Vq2oVO6TGjF0vfIDVaF81ui+EbOsSi56xDghnEEi1Mzh+FQhucxnwP4awzghm0QeMLrNE+dEQwoyfTxsk3et5Kb9k7CWedEMxgWY0DjwMecIVw1gHBjN64Uok301uGcNY4t8uAickd/x8/AGXc+qB2L56DMrTBF9SGCWY0beURRacbo284cn+of6b1oE7C01vGlHDWKMEMoAN3B2fBvJ4Lz8i1zhqUPBQZACCGPz1mKaXsdG5e8/Y472ScWZhyDp/BvOM4pruWk1uOMz0nT9ZjlHX2VLlaef/o5etl3lw3t1drWju1eqi8pdaRU5nQjtsbtMljl74p7O5FAuSYN0ZLV59O28jcxitcUBTMgKlvY7f0TbG1b7td2eiJONWL+fT7XxW9fBSzFraWQtdeENvbLnLC29KybiWYAXNuj9GOt1+JGb18j1sJsL/qqYGe8GkbNCxMb/3fVVu9cLt/nw6FWfo9h2AG/Jg12Gn2NwdAaMc0jPVw64wzPWallrO2zO//F61XwQygTXrLYpfvzbZC0trfzpx2fKrrb2tbKx56BTMA6N+TvWZry7iy/LUes+6+oAhmEJ+eAeBN9nrccsfBluhRe7ztFcwgvh7Gh1BWd70EB0UvX22t18/S4P/p72sXBcw/85WrMqtdISGYAcA7tPwlb23A/x2f5fu+VcKZYAbQFr1lsctHedOgtNaLVmqbyNm+bg1sghkA9KPH4JpWpu/4rDmhbOveZkdPp/4hmAG0Q29Z7PJRxpGxX4/fmX9h+UVPEQtmANCJjBtB3x1g9k7zbf09Dfvle2KM2d7yt3rI5mU6XEbBDNqht+Dd9JbFLh//PB389kLSsTffePrJzv+nz+svPzFFMIPgPg2FA9LL3f1IrOiP3IpePh5xZBtodntJDTzMFADgFf5XuwAAAPwjmAEABCGYAQAEIZgBAAQhmAEABCGYAQAEIZgBAATx/5KdyTAnNxxCAAAAAElFTkSuQmCC" />
			</header>
		</div>
		<div class="container">
			<div class="row">
				<div class="col-md-3 hidden-sm hidden-xs">
					<div id="sidebar" data-spy="affix" data-offset-top="276">
						<h2>API usage</h2>
						<ul class="list-group nav nav-pills nav-stacked">
							<li class="active"><a href="#join">join</a></li>
							<li><a href="#leave">leave</a></li>
							<li><a href="#notice">notice</a></li>
							<li><a href="#privmsg">privmsg</a></li>
						</ul>
						<h2>Channel list</h2>
						<div class="channel-list-view"></div>
					</div>
				</div>
				<div class="col-md-9">
					<div id="channel_list" class="visible-sm visible-xs">
						<h2>Channel list</h2>
						<div class="channel-list-view"></div>
					</div>
					<h2>API usage</h2>
					<div id="join" class="panel panel-success">
						<div class="panel-heading">
							<span class="badge">POST</span>
							{{.}}/join
						</div>
						<div class="panel-body">
							<h3>channel join</h3>
							<form action="{{.}}/join" method="POST" role="form">
								<table class="table table-condensed">
									<thead>
										<tr>
											<th>Parameter</th>
											<th>Value</th>
											<th>Optional</th>
											<th>Description</th>
										</tr>
									</thead>
									<tbody>
										<tr>
											<td>channel</td>
											<td>
												<div class="form-group">
													<input type="text" class="form-control" name="channel" />
												</div>
											</td>
											<td></td>
											<td>チャンネル名</td>
										</tr>
										<tr>
											<td>channel_keyword</td>
											<td>
												<div class="form-group">
													<input type="text" class="form-control" name="channel_keyword" />
												</div>
											</td>
											<td>✓</td>
											<td>パスワード</td>
										</tr>
									</tbody>
								</table>
								<input type="submit" class="btn btn-default btn-block" value="try" />
							</form>
						</div>
					</div>
					<div id="leave" class="panel panel-success">
						<div class="panel-heading">
							<span class="badge">POST</span>
							{{.}}/leave
						</div>
						<div class="panel-body">
							<h3>channel leave</h3>
							<form action="{{.}}/leave" method="POST" role="form">
								<table class="table table-condensed">
									<thead>
										<tr>
											<th>Parameter</th>
											<th>Value</th>
											<th>Optional</th>
											<th>Description</th>
										</tr>
									</thead>
									<tbody>
										<tr>
											<td>channel</td>
											<td>
												<div class="form-group">
													<input type="text" class="form-control" name="channel" />
												</div>
											</td>
											<td></td>
											<td>チャンネル名</td>
										</tr>
									</tbody>
								</table>
								<input type="submit" class="btn btn-default btn-block" value="try" />
							</form>
						</div>
					</div>
					<div id="notice" class="panel panel-success">
						<div class="panel-heading">
							<span class="badge">POST</span>
							{{.}}/notice
						</div>
						<div class="panel-body">
							<h3>sent notie message to channel</h3>
							<form action="{{.}}/notice" method="POST" role="form">
								<table class="table table-condensed">
									<thead>
										<tr>
											<th>Parameter</th>
											<th>Value</th>
											<th>Optional</th>
											<th>Description</th>
										</tr>
									</thead>
									<tbody>
										<tr>
											<td>channel</td>
											<td>
												<div class="form-group">
													<input type="text" class="form-control" name="channel" />
												</div>
											</td>
											<td></td>
											<td>チャンネル名</td>
										</tr>
										<tr>
											<td>message</td>
											<td>
												<div class="form-group">
													<input type="text" class="form-control" name="message" />
												</div>
											</td>
											<td></td>
											<td>メッセージ</td>
										</tr>
									</tbody>
								</table>
								<input type="submit" class="btn btn-default btn-block" value="try" />
							</form>
						</div>
					</div>
					<div id="privmsg" class="panel panel-success">
						<div class="panel-heading">
							<span class="badge">POST</span>
							{{.}}/privmsg
						</div>
						<div class="panel-body">
							<h3>sent privmsg message to channel</h3>
							<form action="{{.}}/privmsg" method="POST" role="form">
								<table class="table table-condensed">
									<thead>
										<tr>
											<th>Parameter</th>
											<th>Value</th>
											<th>Optional</th>
											<th>Description</th>
										</tr>
									</thead>
									<tbody>
										<tr>
											<td>channel</td>
											<td>
												<div class="form-group">
													<input type="text" class="form-control" name="channel" />
												</div>
											</td>
											<td></td>
											<td>チャンネル名</td>
										</tr>
										<tr>
											<td>message</td>
											<td>
												<div class="form-group">
													<input type="text" class="form-control" name="message" />
												</div>
											</td>
											<td></td>
											<td>メッセージ</td>
										</tr>
									</tbody>
								</table>
								<input type="submit" class="btn btn-default btn-block" value="try" />
							</form>
						</div>
					</div>
				</div>
			</div>
		</div>
		<!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
		<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
		<!-- Include all compiled plugins (below), or include individual files as needed -->
		<script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
		<script>
			$.ajax({
				url: "{{.}}/channel_list",
			}).done(function(data){
				$(".channel-list-view").each(function(index){
					arr = data.split(/\r\n|\r|\n/);
					var html = "<ul>"
					for (i = 0; i < arr.length; i++) {
						if (arr[i] != "") {
							html += "<li>" + arr[i] + "</li>"
						}
					}
					html += "</ul>"
					$(this).html(html)
				})
			})
		</script>
	</body>
</html>`))
