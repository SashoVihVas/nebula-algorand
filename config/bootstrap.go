package config

var (

	// BootstrapPeersFilecoin extracted from:
	// https://github.com/filecoin-project/lotus/blob/b691adc4874e5e28353f036c077c08ef00ec3b2b/build/bootstrap/mainnet.pi
	// https://github.com/filecoin-project/lotus/blob/master/build/bootstrap/mainnet.pi
	BootstrapPeersFilecoin = []string{
		"/dns/bootstrap.filecoin.chain.love/tcp/1235/p2p/12D3KooWBF8cpp65hp2u9LK5mh19x67ftAam84z9LsfaquTDSBpt",
		"/dns/bootstrap-venus.mainnet.filincubator.com/tcp/8888/p2p/QmQu8C6deXwKvJP2D8B6QGyhngc3ZiDnFzEHBDx8yeBXST",
		"/dns/bootstrap-mainnet-0.chainsafe-fil.io/tcp/34000/p2p/12D3KooWKKkCZbcigsWTEu1cgNetNbZJqeNtysRtFpq7DTqw3eqH",
		"/dns/bootstrap-mainnet-1.chainsafe-fil.io/tcp/34000/p2p/12D3KooWGnkd9GQKo3apkShQDaq1d6cKJJmsVe6KiQkacUk1T8oZ",
		"/dns/bootstrap-mainnet-2.chainsafe-fil.io/tcp/34000/p2p/12D3KooWHQRSDFv4FvAjtU32shQ7znz7oRbLBryXzZ9NMK2feyyH",
		"/dns/n1.mainnet.fil.devtty.eu/udp/443/quic-v1/p2p/12D3KooWAke3M2ji7tGNKx3BQkTHCyxVhtV1CN68z6Fkrpmfr37F",
		"/dns/n1.mainnet.fil.devtty.eu/tcp/443/p2p/12D3KooWAke3M2ji7tGNKx3BQkTHCyxVhtV1CN68z6Fkrpmfr37F",
		"/dns/n1.mainnet.fil.devtty.eu/udp/443/quic-v1/webtransport/certhash/uEiAWlgd8EqbNhYLv86OdRvXHMosaUWFFDbhgGZgCkcmKnQ/certhash/uEiAvtq6tvZOZf_sIuityDDTyAXDJPfXSRRDK2xy9UVPsqA/p2p/12D3KooWAke3M2ji7tGNKx3BQkTHCyxVhtV1CN68z6Fkrpmfr37F",
	}

	// BootstrapPeersKusama extracted from:
	//   https://github.com/paritytech/polkadot-sdk/blob/master/polkadot/node/service/chain-specs/kusama.json
	BootstrapPeersKusama = []string{
		"/dns/kusama-bootnode-0.polkadot.io/tcp/30333/p2p/12D3KooWSueCPH3puP2PcvqPJdNaDNF3jMZjtJtDiSy35pWrbt5h",
		"/dns/kusama-bootnode-0.polkadot.io/tcp/30334/ws/p2p/12D3KooWSueCPH3puP2PcvqPJdNaDNF3jMZjtJtDiSy35pWrbt5h",
		"/dns/kusama-bootnode-0.polkadot.io/tcp/443/wss/p2p/12D3KooWSueCPH3puP2PcvqPJdNaDNF3jMZjtJtDiSy35pWrbt5h",
		"/dns/kusama-bootnode-1.polkadot.io/tcp/30333/p2p/12D3KooWQKqane1SqWJNWMQkbia9qiMWXkcHtAdfW5eVF8hbwEDw",
		"/dns/kusama-bootnode-1.polkadot.io/tcp/30334/ws/p2p/12D3KooWQKqane1SqWJNWMQkbia9qiMWXkcHtAdfW5eVF8hbwEDw",
		"/dns/kusama-bootnode-1.polkadot.io/tcp/443/wss/p2p/12D3KooWQKqane1SqWJNWMQkbia9qiMWXkcHtAdfW5eVF8hbwEDw",
		"/dns/kusama-boot.dwellir.com/tcp/30333/ws/p2p/12D3KooWFj2ndawdYyk2spc42Y2arYwb2TUoHLHFAsKuHRzWXwoJ",
		"/dns/kusama-boot.dwellir.com/tcp/443/wss/p2p/12D3KooWFj2ndawdYyk2spc42Y2arYwb2TUoHLHFAsKuHRzWXwoJ",
		"/dns/boot.stake.plus/tcp/31333/p2p/12D3KooWLa1UyG5xLPds2GbiRBCTJjpsVwRWHWN7Dff14yiNJRpR",
		"/dns/boot.stake.plus/tcp/31334/wss/p2p/12D3KooWLa1UyG5xLPds2GbiRBCTJjpsVwRWHWN7Dff14yiNJRpR",
		"/dns/boot-node.helikon.io/tcp/7060/p2p/12D3KooWL4KPqfAsPE2aY1g5Zo1CxsDwcdJ7mmAghK7cg6M2fdbD",
		"/dns/boot-node.helikon.io/tcp/7062/wss/p2p/12D3KooWL4KPqfAsPE2aY1g5Zo1CxsDwcdJ7mmAghK7cg6M2fdbD",
		"/dns/kusama.bootnode.amforc.com/tcp/30333/p2p/12D3KooWLx6nsj6Fpd8biP1VDyuCUjazvRiGWyBam8PsqRJkbUb9",
		"/dns/kusama.bootnode.amforc.com/tcp/30334/wss/p2p/12D3KooWLx6nsj6Fpd8biP1VDyuCUjazvRiGWyBam8PsqRJkbUb9",
		"/dns/kusama.bootnodes.polkadotters.com/tcp/30311/p2p/12D3KooWHB5rTeNkQdXNJ9ynvGz8Lpnmsctt7Tvp7mrYv6bcwbPG",
		"/dns/kusama.bootnodes.polkadotters.com/tcp/30313/wss/p2p/12D3KooWHB5rTeNkQdXNJ9ynvGz8Lpnmsctt7Tvp7mrYv6bcwbPG",
		"/dns/boot-cr.gatotech.network/tcp/33200/p2p/12D3KooWRNZXf99BfzQDE1C8YhuBbuy7Sj18UEf7FNpD8egbURYD",
		"/dns/boot-cr.gatotech.network/tcp/35200/wss/p2p/12D3KooWRNZXf99BfzQDE1C8YhuBbuy7Sj18UEf7FNpD8egbURYD",
		"/dns/boot-kusama.metaspan.io/tcp/23012/p2p/12D3KooWE1tq9ZL9AAxMiUBBqy1ENmh5pwfWabnoBPMo8gFPXhn6",
		"/dns/boot-kusama.metaspan.io/tcp/23015/ws/p2p/12D3KooWE1tq9ZL9AAxMiUBBqy1ENmh5pwfWabnoBPMo8gFPXhn6",
		"/dns/boot-kusama.metaspan.io/tcp/23016/wss/p2p/12D3KooWE1tq9ZL9AAxMiUBBqy1ENmh5pwfWabnoBPMo8gFPXhn6",
		"/dns/kusama-bootnode.turboflakes.io/tcp/30305/p2p/12D3KooWR6cMhCYRhbJdqYZfzWZT6bcck3unpRLk8GBQGmHBgPwu",
		"/dns/kusama-bootnode.turboflakes.io/tcp/30405/wss/p2p/12D3KooWR6cMhCYRhbJdqYZfzWZT6bcck3unpRLk8GBQGmHBgPwu",
		"/dns/kusama-boot-ng.dwellir.com/tcp/443/wss/p2p/12D3KooWLswepVYVdCNduvWRTyNTaDMXEBcmvJdZ9Bhw3u2Jhad2",
		"/dns/kusama-boot-ng.dwellir.com/tcp/30334/p2p/12D3KooWLswepVYVdCNduvWRTyNTaDMXEBcmvJdZ9Bhw3u2Jhad2",
		"/dns/kusama-bootnode.radiumblock.com/tcp/30335/wss/p2p/12D3KooWGzKffWe7JSXeKMQeSQC5xfBafZtgBDCuBVxmwe2TJRuc",
		"/dns/kusama-bootnode.radiumblock.com/tcp/30333/p2p/12D3KooWGzKffWe7JSXeKMQeSQC5xfBafZtgBDCuBVxmwe2TJRuc",
		"/dns/ksm-bootnode.stakeworld.io/tcp/30300/p2p/12D3KooWFRin7WWVS6RgUsSpkfUHSv4tfGKnr2zJPmf1pbMv118H",
		"/dns/ksm-bootnode.stakeworld.io/tcp/30301/ws/p2p/12D3KooWFRin7WWVS6RgUsSpkfUHSv4tfGKnr2zJPmf1pbMv118H",
		"/dns/ksm-bootnode.stakeworld.io/tcp/30302/wss/p2p/12D3KooWFRin7WWVS6RgUsSpkfUHSv4tfGKnr2zJPmf1pbMv118H",
		"/dns/ksm14.rotko.net/tcp/35224/wss/p2p/12D3KooWAa5THTw8HPfnhEei23HdL8P9McBXdozG2oTtMMksjZkK",
		"/dns/ksm14.rotko.net/tcp/33224/p2p/12D3KooWAa5THTw8HPfnhEei23HdL8P9McBXdozG2oTtMMksjZkK",
		"/dns/ibp-boot-kusama.luckyfriday.io/tcp/30333/p2p/12D3KooW9vu1GWHBuxyhm7rZgD3fhGZpNajPXFexadvhujWMgwfT",
		"/dns/boot-kusama.luckyfriday.io/tcp/443/wss/p2p/12D3KooWS1Lu6DmK8YHSvkErpxpcXmk14vG6y4KVEFEkd9g62PP8",
		"/dns/ibp-boot-kusama.luckyfriday.io/tcp/30334/wss/p2p/12D3KooW9vu1GWHBuxyhm7rZgD3fhGZpNajPXFexadvhujWMgwfT",
	}

	// BootstrapPeersPolkadot extracted from:
	//   https://github.com/paritytech/polkadot-sdk/blob/master/polkadot/node/service/chain-specs/polkadot.json
	// curl https://raw.githubusercontent.com/paritytech/polkadot-sdk/refs/heads/master/polkadot/node/service/chain-specs/polkadot.json | jq .bootNodes
	BootstrapPeersPolkadot = []string{
		"/dns/polkadot-bootnode-0.polkadot.io/tcp/30333/p2p/12D3KooWSz8r2WyCdsfWHgPyvD8GKQdJ1UAiRmrcrs8sQB3fe2KU",
		"/dns/polkadot-bootnode-0.polkadot.io/tcp/30334/ws/p2p/12D3KooWSz8r2WyCdsfWHgPyvD8GKQdJ1UAiRmrcrs8sQB3fe2KU",
		"/dns/polkadot-bootnode-0.polkadot.io/tcp/443/wss/p2p/12D3KooWSz8r2WyCdsfWHgPyvD8GKQdJ1UAiRmrcrs8sQB3fe2KU",
		"/dns/polkadot-bootnode-1.polkadot.io/tcp/30333/p2p/12D3KooWFN2mhgpkJsDBuNuE5427AcDrsib8EoqGMZmkxWwx3Md4",
		"/dns/polkadot-bootnode-1.polkadot.io/tcp/30334/ws/p2p/12D3KooWFN2mhgpkJsDBuNuE5427AcDrsib8EoqGMZmkxWwx3Md4",
		"/dns/polkadot-bootnode-1.polkadot.io/tcp/443/wss/p2p/12D3KooWFN2mhgpkJsDBuNuE5427AcDrsib8EoqGMZmkxWwx3Md4",
		"/dns/polkadot-boot.dwellir.com/tcp/30334/ws/p2p/12D3KooWKvdDyRKqUfSAaUCbYiLwKY8uK3wDWpCuy2FiDLbkPTDJ",
		"/dns/polkadot-boot.dwellir.com/tcp/443/wss/p2p/12D3KooWKvdDyRKqUfSAaUCbYiLwKY8uK3wDWpCuy2FiDLbkPTDJ",
		"/dns/polkadot.boot.stake.plus/tcp/30334/wss/p2p/12D3KooWCZKEvAMJRk9nwTHJcTjgVw6bDEqryQ3B7n7scNtfNqPB",
		"/dns/polkadot.boot.stake.plus/tcp/31334/wss/p2p/12D3KooWMFwJV935CyJXE8twfkKxRDnNWeEFd8jZWaoWZF22Hv8S",
		"/dns/boot-node.helikon.io/tcp/7070/p2p/12D3KooWS9ZcvRxyzrSf6p63QfTCWs12nLoNKhGux865crgxVA4H",
		"/dns/boot-node.helikon.io/tcp/7072/wss/p2p/12D3KooWS9ZcvRxyzrSf6p63QfTCWs12nLoNKhGux865crgxVA4H",
		"/dns/polkadot.bootnode.amforc.com/tcp/30001/p2p/12D3KooWT2HyZx5C6BBeLbCKhYG2SqJYuiu7sLMxGzUcQBko3BMr",
		"/dns/polkadot.bootnode.amforc.com/tcp/29999/wss/p2p/12D3KooWT2HyZx5C6BBeLbCKhYG2SqJYuiu7sLMxGzUcQBko3BMr",
		"/dns/polkadot.bootnodes.polkadotters.com/tcp/30314/p2p/12D3KooWPAVUgBaBk6n8SztLrMk8ESByncbAfRKUdxY1nygb9zG3",
		"/dns/polkadot.bootnodes.polkadotters.com/tcp/30316/wss/p2p/12D3KooWPAVUgBaBk6n8SztLrMk8ESByncbAfRKUdxY1nygb9zG3",
		"/dns/boot.gatotech.network/tcp/33100/p2p/12D3KooWK4E16jKk9nRhvC4RfrDVgcZzExg8Q3Q2G7ABUUitks1w",
		"/dns/boot.gatotech.network/tcp/35100/wss/p2p/12D3KooWK4E16jKk9nRhvC4RfrDVgcZzExg8Q3Q2G7ABUUitks1w",
		"/dns/boot.metaspan.io/tcp/13012/p2p/12D3KooWRjHFApinuqSBjoaDjQHvxwubQSpEVy5hrgC9Smvh92WF",
		"/dns/boot.metaspan.io/tcp/13015/ws/p2p/12D3KooWRjHFApinuqSBjoaDjQHvxwubQSpEVy5hrgC9Smvh92WF",
		"/dns/boot.metaspan.io/tcp/13016/wss/p2p/12D3KooWRjHFApinuqSBjoaDjQHvxwubQSpEVy5hrgC9Smvh92WF",
		"/dns/polkadot-bootnode.turboflakes.io/tcp/30300/p2p/12D3KooWHJBMZgt7ymAdTRtadPcGXpJw79vBGe8z53r9JMkZW7Ha",
		"/dns/polkadot-bootnode.turboflakes.io/tcp/30400/wss/p2p/12D3KooWHJBMZgt7ymAdTRtadPcGXpJw79vBGe8z53r9JMkZW7Ha",
		"/dns/polkadot-boot-ng.dwellir.com/tcp/443/wss/p2p/12D3KooWFFqjBKoSdQniRpw1Y8W6kkV7takWv1DU2ZMkaA81PYVq",
		"/dns/polkadot-boot-ng.dwellir.com/tcp/30336/p2p/12D3KooWFFqjBKoSdQniRpw1Y8W6kkV7takWv1DU2ZMkaA81PYVq",
		"/dns/polkadot-bootnode.radiumblock.com/tcp/30335/wss/p2p/12D3KooWNwWNRrPrTk4qMah1YszudMjxNw2qag7Kunhw3Ghs9ea5",
		"/dns/polkadot-bootnode.radiumblock.com/tcp/30333/p2p/12D3KooWNwWNRrPrTk4qMah1YszudMjxNw2qag7Kunhw3Ghs9ea5",
		"/dns/dot-bootnode.stakeworld.io/tcp/30310/p2p/12D3KooWAb5MyC1UJiEQJk4Hg4B2Vi3AJdqSUhTGYUqSnEqCFMFg",
		"/dns/dot-bootnode.stakeworld.io/tcp/30311/ws/p2p/12D3KooWAb5MyC1UJiEQJk4Hg4B2Vi3AJdqSUhTGYUqSnEqCFMFg",
		"/dns/dot-bootnode.stakeworld.io/tcp/30312/wss/p2p/12D3KooWAb5MyC1UJiEQJk4Hg4B2Vi3AJdqSUhTGYUqSnEqCFMFg",
		"/dns/polkadot.boot.rotko.net/tcp/31001/p2p/12D3KooWPyEvPEXghnMC67Gff6PuZiSvfx3fmziKiPZcGStZ5xff",
		"/dns/polkadot.boot.rotko.net/tcp/30335/wss/p2p/12D3KooWPyEvPEXghnMC67Gff6PuZiSvfx3fmziKiPZcGStZ5xff",
		"/dns/ibp-boot-polkadot.luckyfriday.io/tcp/30333/p2p/12D3KooWEjk6QXrZJ26fLpaajisJGHiz6WiQsR8k7mkM9GmWKnRZ",
		"/dns/ibp-boot-polkadot.luckyfriday.io/tcp/30334/wss/p2p/12D3KooWEjk6QXrZJ26fLpaajisJGHiz6WiQsR8k7mkM9GmWKnRZ",
		"/dns/boot-polkadot.luckyfriday.io/tcp/443/wss/p2p/12D3KooWAdyiVAaeGdtBt6vn5zVetwA4z4qfm9Fi2QCSykN1wTBJ",
		"/dns4/polkadot-0.boot.onfinality.io/tcp/24446/ws/p2p/12D3KooWT1PWaNdAwYrSr89dvStnoGdH3t4LNRbcVNN4JCtsotkR",
		"/dns/polkadot.bootnode.stkd.io/tcp/30633/wss/p2p/12D3KooWEymrFRHz6c17YP3FAyd8kXS5gMRLgkW4U77ZJD2ZNCLZ",
	}

	// BootstrapPeersRococo extracted from:
	//   https://github.com/paritytech/polkadot-sdk/blob/master/polkadot/node/service/chain-specs/rococo.json
	BootstrapPeersRococo = []string{
		"/dns/rococo-bootnode-0.polkadot.io/tcp/30333/p2p/12D3KooWGikJMBmRiG5ofCqn8aijCijgfmZR5H9f53yUF3srm6Nm",
		"/dns/rococo-bootnode-1.polkadot.io/tcp/30333/p2p/12D3KooWLDfH9mHRCidrd5NfQjp7rRMUcJSEUwSvEKyu7xU2cG3d",
		"/dns/rococo-bootnode-0.polkadot.io/tcp/30334/ws/p2p/12D3KooWGikJMBmRiG5ofCqn8aijCijgfmZR5H9f53yUF3srm6Nm",
		"/dns/rococo-bootnode-1.polkadot.io/tcp/30334/ws/p2p/12D3KooWLDfH9mHRCidrd5NfQjp7rRMUcJSEUwSvEKyu7xU2cG3d",
		"/dns/rococo-bootnode-0.polkadot.io/tcp/443/wss/p2p/12D3KooWGikJMBmRiG5ofCqn8aijCijgfmZR5H9f53yUF3srm6Nm",
		"/dns/rococo-bootnode-1.polkadot.io/tcp/443/wss/p2p/12D3KooWLDfH9mHRCidrd5NfQjp7rRMUcJSEUwSvEKyu7xU2cG3d",
	}

	// BootstrapPeersWestend extracted from:
	//   https://github.com/paritytech/polkadot-sdk/blob/master/polkadot/node/service/chain-specs/westend.json
	BootstrapPeersWestend = []string{
		"/dns/westend-bootnode-0.polkadot.io/tcp/30333/p2p/12D3KooWKer94o1REDPtAhjtYR4SdLehnSrN8PEhBnZm5NBoCrMC",
		"/dns/westend-bootnode-0.polkadot.io/tcp/30334/ws/p2p/12D3KooWKer94o1REDPtAhjtYR4SdLehnSrN8PEhBnZm5NBoCrMC",
		"/dns/westend-bootnode-0.polkadot.io/tcp/443/wss/p2p/12D3KooWKer94o1REDPtAhjtYR4SdLehnSrN8PEhBnZm5NBoCrMC",
		"/dns/westend-bootnode-1.polkadot.io/tcp/30333/p2p/12D3KooWPVPzs42GvRBShdUMtFsk4SvnByrSdWqb6aeAAHvLMSLS",
		"/dns/westend-bootnode-1.polkadot.io/tcp/30334/ws/p2p/12D3KooWPVPzs42GvRBShdUMtFsk4SvnByrSdWqb6aeAAHvLMSLS",
		"/dns/westend-bootnode-1.polkadot.io/tcp/443/wss/p2p/12D3KooWPVPzs42GvRBShdUMtFsk4SvnByrSdWqb6aeAAHvLMSLS",
		"/dns/boot.stake.plus/tcp/32333/p2p/12D3KooWK8fjVoSvMq5copQYMsdYreSGPGgcMbGMgbMDPfpf3sm7",
		"/dns/boot.stake.plus/tcp/32334/wss/p2p/12D3KooWK8fjVoSvMq5copQYMsdYreSGPGgcMbGMgbMDPfpf3sm7",
		"/dns/boot-node.helikon.io/tcp/7080/p2p/12D3KooWRFDPyT8vA8mLzh6dJoyujn4QNjeqi6Ch79eSMz9beKXC",
		"/dns/boot-node.helikon.io/tcp/7082/wss/p2p/12D3KooWRFDPyT8vA8mLzh6dJoyujn4QNjeqi6Ch79eSMz9beKXC",
		"/dns/westend.bootnode.amforc.com/tcp/30333/p2p/12D3KooWJ5y9ZgVepBQNW4aabrxgmnrApdVnscqgKWiUu4BNJbC8",
		"/dns/westend.bootnode.amforc.com/tcp/30334/wss/p2p/12D3KooWJ5y9ZgVepBQNW4aabrxgmnrApdVnscqgKWiUu4BNJbC8",
		"/dns/westend.bootnodes.polkadotters.com/tcp/30308/p2p/12D3KooWHPHb64jXMtSRJDrYFATWeLnvChL8NtWVttY67DCH1eC5",
		"/dns/westend.bootnodes.polkadotters.com/tcp/30310/wss/p2p/12D3KooWHPHb64jXMtSRJDrYFATWeLnvChL8NtWVttY67DCH1eC5",
		"/dns/boot-cr.gatotech.network/tcp/33300/p2p/12D3KooWQGR1vUhoy6mvQorFp3bZFn6NNezhQZ6NWnVV7tpFgoPd",
		"/dns/boot-cr.gatotech.network/tcp/35300/wss/p2p/12D3KooWQGR1vUhoy6mvQorFp3bZFn6NNezhQZ6NWnVV7tpFgoPd",
		"/dns/boot-westend.metaspan.io/tcp/33012/p2p/12D3KooWNTau7iG4G9cUJSwwt2QJP1W88pUf2SgqsHjRU2RL8pfa",
		"/dns/boot-westend.metaspan.io/tcp/33015/ws/p2p/12D3KooWNTau7iG4G9cUJSwwt2QJP1W88pUf2SgqsHjRU2RL8pfa",
		"/dns/boot-westend.metaspan.io/tcp/33016/wss/p2p/12D3KooWNTau7iG4G9cUJSwwt2QJP1W88pUf2SgqsHjRU2RL8pfa",
		"/dns/westend-bootnode.turboflakes.io/tcp/30310/p2p/12D3KooWJvPDCZmReU46ghpCMJCPVUvUCav4WQdKtXQhZgJdH6tZ",
		"/dns/westend-bootnode.turboflakes.io/tcp/30410/wss/p2p/12D3KooWJvPDCZmReU46ghpCMJCPVUvUCav4WQdKtXQhZgJdH6tZ",
		"/dns/westend-boot-ng.dwellir.com/tcp/443/wss/p2p/12D3KooWJifoDhCL3swAKt7MWhFb7wLRFD9oG33AL3nAathmU24x",
		"/dns/westend-boot-ng.dwellir.com/tcp/30335/p2p/12D3KooWJifoDhCL3swAKt7MWhFb7wLRFD9oG33AL3nAathmU24x",
		"/dns/westend-bootnode.radiumblock.com/tcp/30335/wss/p2p/12D3KooWJBowJuX1TaWNWHt8Dz8z44BoCZunLCfFqxA2rLTn6TBD",
		"/dns/westend-bootnode.radiumblock.com/tcp/30333/p2p/12D3KooWJBowJuX1TaWNWHt8Dz8z44BoCZunLCfFqxA2rLTn6TBD",
		"/dns/wnd-bootnode.stakeworld.io/tcp/30320/p2p/12D3KooWBYdKipcNbrV5rCbgT5hco8HMLME7cE9hHC3ckqCKDuzP",
		"/dns/wnd-bootnode.stakeworld.io/tcp/30321/ws/p2p/12D3KooWBYdKipcNbrV5rCbgT5hco8HMLME7cE9hHC3ckqCKDuzP",
		"/dns/wnd-bootnode.stakeworld.io/tcp/30322/wss/p2p/12D3KooWBYdKipcNbrV5rCbgT5hco8HMLME7cE9hHC3ckqCKDuzP",
		"/dns/wnd14.rotko.net/tcp/35234/wss/p2p/12D3KooWLK8Zj1uZ46phU3vQwiDVda8tB76S8J26rXZQLHpwWkDJ",
		"/dns/wnd14.rotko.net/tcp/33234/p2p/12D3KooWLK8Zj1uZ46phU3vQwiDVda8tB76S8J26rXZQLHpwWkDJ",
		"/dns/ibp-boot-westend.luckyfriday.io/tcp/30333/p2p/12D3KooWDg1YEytdwFFNWroFj6gio4YFsMB3miSbHKgdpJteUMB9",
		"/dns/ibp-boot-westend.luckyfriday.io/tcp/30334/wss/p2p/12D3KooWDg1YEytdwFFNWroFj6gio4YFsMB3miSbHKgdpJteUMB9",
	}

	// BootstrapPeersCelestia extracted from:
	//   https://github.com/celestiaorg/celestia-node/blob/3fde8c9acb90bb9d4d1abbdc2c5917d7f3693388/nodebuilder/p2p/bootstrap.go#L39
	BootstrapPeersCelestia = []string{
		"/dns4/da-bridge-1.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWSqZaLcn5Guypo2mrHr297YPJnV8KMEMXNjs3qAS8msw8",
		"/dns4/da-bridge-2.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWQpuTFELgsUypqp9N4a1rKBccmrmQVY8Em9yhqppTJcXf",
		"/dns4/da-bridge-3.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWSGa4huD6ts816navn7KFYiStBiy5LrBQH1HuEahk4TzQ",
		"/dns4/da-bridge-4.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWHBXCmXaUNat6ooynXG837JXPsZpSTeSzZx6DpgNatMmR",
		"/dns4/da-bridge-5.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWDGTBK1a2Ru1qmnnRwP6Dmc44Zpsxi3xbgFk7ATEPfmEU",
		"/dns4/da-bridge-6.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWLTUFyf3QEGqYkHWQS2yCtuUcL78vnKBdXU5gABM1YDeH",
		"/dns4/da-full-1.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWKZCMcwGCYbL18iuw3YVpAZoyb1VBGbx9Kapsjw3soZgr",
		"/dns4/da-full-2.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWE3fmRtHgfk9DCuQFfY3H3JYEnTU3xZozv1Xmo8KWrWbK",
		"/dns4/da-full-3.celestia-bootstrap.net/tcp/2121/p2p/12D3KooWK6Ftsd4XsWCsQZgZPNhTrE5urwmkoo5P61tGvnKmNVyv",
	}

	// BootstrapPeersArabica extracted from:
	//   https://github.com/celestiaorg/celestia-node/blob/3fde8c9acb90bb9d4d1abbdc2c5917d7f3693388/nodebuilder/p2p/bootstrap.go#L51
	BootstrapPeersArabica = []string{
		"/dns4/da-bridge.celestia-arabica-10.com/tcp/2121/p2p/12D3KooWM3e9MWtyc8GkP8QRt74Riu17QuhGfZMytB2vq5NwkWAu",
		"/dns4/da-bridge-2.celestia-arabica-10.com/tcp/2121/p2p/12D3KooWKj8mcdiBGxQRe1jqhaMnh2tGoC3rPDmr5UH2q8H4WA9M",
		"/dns4/da-full-1.celestia-arabica-10.com/tcp/2121/p2p/12D3KooWBWkgmN7kmJSFovVrCjkeG47FkLGq7yEwJ2kEqNKCsBYk",
		"/dns4/da-full-2.celestia-arabica-10.com/tcp/2121/p2p/12D3KooWRByRF67a2kVM2j4MP5Po3jgTw7H2iL2Spu8aUwPkrRfP",
	}

	// BootstrapPeersMocha extracted from:
	//   https://github.com/celestiaorg/celestia-node/blob/3fde8c9acb90bb9d4d1abbdc2c5917d7f3693388/nodebuilder/p2p/bootstrap.go#L57
	BootstrapPeersMocha = []string{
		"/dns4/da-bridge-mocha-4.celestia-mocha.com/tcp/2121/p2p/12D3KooWCBAbQbJSpCpCGKzqz3rAN4ixYbc63K68zJg9aisuAajg",
		"/dns4/da-bridge-mocha-4-2.celestia-mocha.com/tcp/2121/p2p/12D3KooWK6wJkScGQniymdWtBwBuU36n6BRXp9rCDDUD6P5gJr3G",
		"/dns4/da-full-1-mocha-4.celestia-mocha.com/tcp/2121/p2p/12D3KooWCUHPLqQXZzpTx1x3TAsdn3vYmTNDhzg66yG8hqoxGGN8",
		"/dns4/da-full-2-mocha-4.celestia-mocha.com/tcp/2121/p2p/12D3KooWR6SHsXPkkvhCRn6vp1RqSefgaT1X1nMNvrVjU2o3GoYy",
	}

	// BootstrapPeersBlockspaceRace extracted from:
	//   https://github.com/celestiaorg/celestia-node/blob/9c0a5fb0626ada6e6cdb8bcd816d01a3aa5043ad/nodebuilder/p2p/bootstrap.go#L39
	BootstrapPeersBlockspaceRace = []string{
		"/dns4/bootstr-incent-3.celestia.tools/tcp/2121/p2p/12D3KooWNzdKcHagtvvr6qtjcPTAdCN6ZBiBLH8FBHbihxqu4GZx",
		"/dns4/bootstr-incent-2.celestia.tools/tcp/2121/p2p/12D3KooWNJZyWeCsrKxKrxsNM1RVL2Edp77svvt7Cosa63TggC9m",
		"/dns4/bootstr-incent-1.celestia.tools/tcp/2121/p2p/12D3KooWBtxdBzToQwnS4ySGpph9PtGmmjEyATkgX3PfhAo4xmf7",
	}

	// BootstrapPeersEthereumConsensus extracted from:
	//   https://github.com/eth-clients/mainnet/blob/main/metadata/bootstrap_nodes.yaml
	BootstrapPeersEthereumConsensus = []string{
		// Teku team's bootnodes
		"enr:-Iu4QLm7bZGdAt9NSeJG0cEnJohWcQTQaI9wFLu3Q7eHIDfrI4cwtzvEW3F3VbG9XdFXlrHyFGeXPn9snTCQJ9bnMRABgmlkgnY0gmlwhAOTJQCJc2VjcDI1NmsxoQIZdZD6tDYpkpEfVo5bgiU8MGRjhcOmHGD2nErK0UKRrIN0Y3CCIyiDdWRwgiMo",
		"enr:-Iu4QEDJ4Wa_UQNbK8Ay1hFEkXvd8psolVK6OhfTL9irqz3nbXxxWyKwEplPfkju4zduVQj6mMhUCm9R2Lc4YM5jPcIBgmlkgnY0gmlwhANrfESJc2VjcDI1NmsxoQJCYz2-nsqFpeEj6eov9HSi9QssIVIVNr0I89J1vXM9foN0Y3CCIyiDdWRwgiMo",
		// Prylab team's bootnodes
		"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
		"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
		"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
		// Lighthouse team's bootnodes
		"enr:-Le4QPUXJS2BTORXxyx2Ia-9ae4YqA_JWX3ssj4E_J-3z1A-HmFGrU8BpvpqhNabayXeOZ2Nq_sbeDgtzMJpLLnXFgAChGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISsaa0Zg2lwNpAkAIkHAAAAAPA8kv_-awoTiXNlY3AyNTZrMaEDHAD2JKYevx89W0CcFJFiskdcEzkH_Wdv9iW42qLK79ODdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QLHZDSvkLfqgEo8IWGG96h6mxwe_PsggC20CL3neLBjfXLGAQFOPSltZ7oP6ol54OvaNqO02Rnvb8YmDR274uq8ChGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISLosQxg2lwNpAqAX4AAAAAAPA8kv_-ax65iXNlY3AyNTZrMaEDBJj7_dLFACaxBfaI8KZTh_SSJUjhyAyfshimvSqo22WDdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QH6LQrusDbAHPjU_HcKOuMeXfdEB5NJyXgHWFadfHgiySqeDyusQMvfphdYWOzuSZO9Uq2AMRJR5O4ip7OvVma8BhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISLY9ncg2lwNpAkAh8AgQIBAAAAAAAAAAmXiXNlY3AyNTZrMaECDYCZTZEksF-kmgPholqgVt8IXr-8L7Nu7YrZ7HUpgxmDdWRwgiMohHVkcDaCI4I",
		"enr:-Le4QIqLuWybHNONr933Lk0dcMmAB5WgvGKRyDihy1wHDIVlNuuztX62W51voT4I8qD34GcTEOTmag1bcdZ_8aaT4NUBhGV0aDKQtTA_KgEAAAAAIgEAAAAAAIJpZIJ2NIJpcISLY04ng2lwNpAkAh8AgAIBAAAAAAAAAA-fiXNlY3AyNTZrMaEDscnRV6n1m-D9ID5UsURk0jsoKNXt1TIrj8uKOGW6iluDdWRwgiMohHVkcDaCI4I",
		// EF bootnodes
		"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
		"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
		"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
		"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
		// Nimbus team's bootnodes
		"enr:-LK4QA8FfhaAjlb_BXsXxSfiysR7R52Nhi9JBt4F8SPssu8hdE1BXQQEtVDC3qStCW60LSO7hEsVHv5zm8_6Vnjhcn0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAN4aBKJc2VjcDI1NmsxoQJerDhsJ-KxZ8sHySMOCmTO6sHM3iCFQ6VMvLTe948MyYN0Y3CCI4yDdWRwgiOM",
		"enr:-LK4QKWrXTpV9T78hNG6s8AM6IO4XH9kFT91uZtFg1GcsJ6dKovDOr1jtAAFPnS2lvNltkOGA9k29BUN7lFh_sjuc9QBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhANAdd-Jc2VjcDI1NmsxoQLQa6ai7y9PMN5hpLe5HmiJSlYzMuzP7ZhwRiwHvqNXdoN0Y3CCI4yDdWRwgiOM",
		// Lodestar team's bootnodes
		"enr:-IS4QPi-onjNsT5xAIAenhCGTDl4z-4UOR25Uq-3TmG4V3kwB9ljLTb_Kp1wdjHNj-H8VVLRBSSWVZo3GUe3z6k0E-IBgmlkgnY0gmlwhKB3_qGJc2VjcDI1NmsxoQMvAfgB4cJXvvXeM6WbCG86CstbSxbQBSGx31FAwVtOTYN1ZHCCIyg",
		"enr:-KG4QCb8NC3gEM3I0okStV5BPX7Bg6ZXTYCzzbYyEXUPGcZtHmvQtiJH4C4F2jG7azTcb9pN3JlgpfxAnRVFzJ3-LykBgmlkgnY0gmlwhFPlR9KDaXA2kP6AAAAAAAAAAlBW__4my5iJc2VjcDI1NmsxoQLdUv9Eo9sxCt0tc_CheLOWnX59yHJtkBSOL7kpxdJ6GYN1ZHCCIyiEdWRwNoIjKA",
	}

	// BootstrapPeersEthereumExecution extracted from:
	//   https://github.com/ethereum/go-ethereum/blob/f04e5bde7487ce554930187e766164b18c37d867/params/bootnodes.go#L23
	BootstrapPeersEthereumExecution = []string{
		"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303", // bootnode-aws-ap-southeast-1-001
		"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",   // bootnode-aws-us-east-1-001
		"enode://2b252ab6a1d0f971d9722cb839a42cb81db019ba44c08754628ab4a823487071b5695317c8ccd085219c3a03af063495b2f1da8d18218da2d6a82981b45e6ffc@65.108.70.101:30303", // bootnode-hetzner-hel
		"enode://4aeb4ab6c14b23e2c4cfdce879c04b0748a20d8e9b59e25ded2a08143e265c6c25936e74cbc8e641e3312ca288673d91f2f93f8e277de3cfa444ecdaaf982052@157.90.35.166:30303", // bootnode-hetzner-fsn
	}

	// BootstrapPeersHolesky extracted from:
	//   https://github.com/eth-clients/holesky/blob/main/custom_config_data/bootstrap_nodes.txt
	BootstrapPeersHolesky = []string{
		"enr:-Ku4QFo-9q73SspYI8cac_4kTX7yF800VXqJW4Lj3HkIkb5CMqFLxciNHePmMt4XdJzHvhrCC5ADI4D_GkAsxGJRLnQBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAhnTT-AQFwAP__________gmlkgnY0gmlwhLKAiOmJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyk",
		"enr:-Ku4QPG7F72mbKx3gEQEx07wpYYusGDh-ni6SNkLvOS-hhN-BxIggN7tKlmalb0L5JPoAfqD-akTZ-gX06hFeBEz4WoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAhnTT-AQFwAP__________gmlkgnY0gmlwhJK-DYCJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyk",
		"enr:-LK4QPxe-mDiSOtEB_Y82ozvxn9aQM07Ui8A-vQHNgYGMMthfsfOabaaTHhhJHFCBQQVRjBww_A5bM1rf8MlkJU_l68Eh2F0dG5ldHOIAADAAAAAAACEZXRoMpBpt9l0BAFwAAABAAAAAAAAgmlkgnY0gmlwhLKAiOmJc2VjcDI1NmsxoQJu6T9pclPObAzEVQ53DpVQqjadmVxdTLL-J3h9NFoCeIN0Y3CCIyiDdWRwgiMo",
		"enr:-Ly4QGbOw4xNel5EhmDsJJ-QhC9XycWtsetnWoZ0uRy381GHdHsNHJiCwDTOkb3S1Ade0SFQkWJX_pgb3g8Jfh93rvMBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpBpt9l0BAFwAAABAAAAAAAAgmlkgnY0gmlwhJK-DYCJc2VjcDI1NmsxoQOxKv9sv3zKF8GDewgFGGHKP5HCZZpPpTrwl9eXKAWGxIhzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA",
		"enr:-LS4QG0uV4qvcpJ-HFDJRGBmnlD3TJo7yc4jwK8iP7iKaTlfQ5kZvIDspLMJhk7j9KapuL9yyHaZmwTEZqr10k9XumyCEcmHYXR0bmV0c4gAAAAABgAAAIRldGgykGm32XQEAXAAAAEAAAAAAACCaWSCdjSCaXCErK4j-YlzZWNwMjU2azGhAgfWRBEJlb7gAhXIB5ePmjj2b8io0UpEenq1Kl9cxStJg3RjcIIjKIN1ZHCCIyg",
		"enr:-Le4QLoE1wFHSlGcm48a9ZESb_MRLqPPu6G0vHqu4MaUcQNDHS69tsy-zkN0K6pglyzX8m24mkb-LtBcbjAYdP1uxm4BhGV0aDKQabfZdAQBcAAAAQAAAAAAAIJpZIJ2NIJpcIQ5gR6Wg2lwNpAgAUHQBwEQAAAAAAAAADR-iXNlY3AyNTZrMaEDPMSNdcL92uNIyCsS177Z6KTXlbZakQqxv3aQcWawNXeDdWRwgiMohHVkcDaCI4I",
	}

	// BootstrapPeersAvailMainnetFullNode extracted from:
	// 	https://raw.githubusercontent.com/availproject/avail/main/misc/genesis/mainnet.chain.spec.json
	BootstrapPeersAvailMainnetFullNode = []string{
		"/dns/bootnode-mainnet-001.avail.so/tcp/30333/p2p/12D3KooWBXk3rcfKkvd1YbJ8fHPH4WZy34QCw8Czrvqf6cbmrdKh",
		"/dns/bootnode-mainnet-002.avail.so/tcp/30333/p2p/12D3KooWMJt1z4ap2UacerW7vJprr7Mqws5sqmJXZqg8XpZswiwF",
		"/dns/bootnode-mainnet-003.avail.so/tcp/30333/p2p/12D3KooWK85zotBy9jjgVhHzsvntAWWzRxTosF3RxwnJ9zJsGQbi",
		"/dns/bootnode-mainnet-004.avail.so/tcp/30333/p2p/12D3KooWFyPvzGbTH7qbB3foBDDfwtn4LJsb4ryvTGGCokUyYDm7",
		"/dns/bootnode-mainnet-005.avail.so/tcp/30333/p2p/12D3KooWFB6wBdwS1aacu3QZdmZb2sPDiVxP9iLFecLBiHKZELCX",
		"/dns/bootnode-mainnet-006.avail.so/tcp/30333/p2p/12D3KooWLVJtN3hpUJYPjQzGKiHnbw6fpD5vvH5EkztdP1VAAgrj",
		"/dns/bootnode-mainnet-007.avail.so/tcp/30333/p2p/12D3KooWPfsS6gAYoEiB8EaWbHdrbd6ugcru2kGguz56B9PxD4HT",
		"/dns/bootnode-mainnet-008.avail.so/tcp/30333/p2p/12D3KooWSYiWfRCigRwWU9UzVJfZRz3SkCpYySH9w5ZydXUJWVuJ",
		"/dns/bootnode-mainnet-009.avail.so/tcp/30333/p2p/12D3KooWH1dPXxYyxwYjhbrjJ3PUJht9DyQkgdiJyiTtKVx5Voh4",
		"/dns/bootnode-mainnet-010.avail.so/tcp/30333/p2p/12D3KooWLeXFyp1Ghm7oCt1EghhT39wnsNhZWT1jYEWTHn2pHU3E",
	}
	// BootstrapPeersAvailMainnetLightClient
	BootstrapPeersAvailMainnetLightClient = []string{
		"/dns/bootnode.1.lightclient.mainnet.avail.so/tcp/37000/p2p/12D3KooW9x9qnoXhkHAjdNFu92kMvBRSiFBMAoC5NnifgzXjsuiM",
	}

	// BootstrapPeersAvailTuringLightClient
	BootstrapPeersAvailTuringLightClient = []string{
		"/dns/bootnode.1.lightclient.turing.avail.so/tcp/37000/p2p/12D3KooWBkLsNGaD3SpMaRWtAmWVuiZg1afdNSPbtJ8M8r9ArGRT",
	}

	// BootstrapPeersAvailTurin network extracted from:
	//   https://raw.githubusercontent.com/availproject/avail/main/misc/genesis/testnet.turing.chain.spec.raw.json
	BootstrapPeersAvailTuringFullNode = []string{
		"/dns/bootnode-turing-001.avail.so/tcp/30333/p2p/12D3KooWRYCs192er73R2oyE4QWkKdtoRk1fL28tR3immzxTQQAq",
		"/dns/bootnode-turing-002.avail.so/tcp/30333/p2p/12D3KooWLxzYHmGXpRnN4rAKkkZVzFgpGuqKofPggUnWz4iUR5yT",
		"/dns/bootnode-turing-003.avail.so/tcp/30333/p2p/12D3KooWR8yPHmBZxHPoZrjjvCZEq5mJh9h7yewgM4zGgUVEANcG",
		"/dns/bootnode-turing-004.avail.so/tcp/30333/p2p/12D3KooWKkLgiqxweQ1QCbJrDKeg5HCktZMM1kd81F4jMMTnBtD6",
		"/dns/bootnode-turing-005.avail.so/tcp/30333/p2p/12D3KooWQRTgqX2j2HMideuPKUg12yeCbbp6cCi6xtf4cACYSiKb",
		"/dns/bootnode-turing-006.avail.so/tcp/30333/p2p/12D3KooWPNmovc7n6mET56bmpB8tM6XtjcRb8qBir3sTBWk5cttK",
		"/dns/bootnode-turing-007.avail.so/tcp/30333/p2p/12D3KooWFzqtue91MrdmqVRddA48dPYQVAFBtcFkf5HrmgiqUJ6F",
		"/dns/bootnode-turing-008.avail.so/tcp/30333/p2p/12D3KooWMHRgZhDMUbN55NQcXB19Ww6SnL8shz1un3K6X3K2XoAh",
		"/dns/bootnode-turing-009.avail.so/tcp/30333/p2p/12D3KooWQLdPYPEHRGAJ4J6dGYoiKyjm36VhaSCJkvKAuJBohJYi",
	}

	// BootstrapPeersPactusFullNode extracted from:
	//   https://github.com/pactus-project/pactus/blob/main/config/bootstrap.json
	BootstrapPeersPactusFullNode = []string{
		"/dns/bootstrap1.pactus.org/tcp/21888/p2p/12D3KooWMnDsu8TDTk2VV8uD8zsNSB6eUkqtQs6ttg4bHq9zNaBe",
		"/dns/bootstrap2.pactus.org/tcp/21888/p2p/12D3KooWM39ag7ghta49qybf7McADgT8FLakTYkCsiBvwdnjuG5q",
		"/dns/bootstrap3.pactus.org/tcp/21888/p2p/12D3KooWBCPSZWheet6tMoHbVBCDfBwQm4yzCwcQ8hJ6NMCN97sj",
		"/dns/bootstrap4.pactus.org/tcp/21888/p2p/12D3KooWKg6aLa77yAaqMCb45aH5iQuTr5GzRUWUCJ1sZYB5vnoL",
		"/dns/pactus-bootstrap.sensifai.com/tcp/21888/p2p/12D3KooWLpgBQt7A1W8FhTLVYwQMEnTBEX7uudLvd1YMZsDHqqWh",
		"/ip4/65.108.211.187/tcp/21888/p2p/12D3KooWGLrSiE5KWoc9GpXfryvKqheHj6j78ryMDDC17SY4YoCi",
		"/ip4/84.247.165.249/tcp/21888/p2p/12D3KooWQmv2FcNQfh1EhA98twt8ePdkQaxEPeYfinEYyVS16juY",
		"/ip4/62.171.130.196/tcp/21888/p2p/12D3KooWSkwB8AiWz3VGyG5cgC3ry5dmCw3TthQNc1jvRNe9EEKH",
		"/ip4/135.181.42.222/tcp/21888/p2p/12D3KooWN46UYQNoZvzd9GiJPQC7LByj6RAg8rRxdUTiQ24vHsN1",
		"/dns/pactus-bootstrap.stakers.world/tcp/21888/p2p/12D3KooWK1DAcxf8RJA2PVkwr9ehk1zgXcnjLyPCsYCpvw1NSfhc",
		"/ip4/158.220.91.129/tcp/21888/p2p/12D3KooWGfehwyZJBu6WUMLyzJ8sDJachvhxhqw5Wh4uU3H99GLq",
		"/dns/pactus-bootstrap.ionode.online/tcp/21888/p2p/12D3KooWG5tvvfS8PRd2Bgyxxqrok7QfohkHSprFPxmT3HTh8dVD",
		"/ip4/159.148.146.149/tcp/21888/p2p/12D3KooWKCokWtpdudxgsLRoFcnrW35vhn6w632iGWCga7E5e68Q",
		"/dns/pactus-bootnode.stake.works/tcp/21888/p2p/12D3KooWJPuhW7Ejmb8KprpVNRzn1eUQXDHU3n5eTo2vodgvVxoJ",
		"/ip4/65.108.68.214/tcp/21888/p2p/12D3KooWD7X8BDqV9r6zGmoo2dpMYo4Z6YoyDQp6j4UnB2Bj3Fq2",
		"/ip4/65.21.180.80/tcp/21888/p2p/12D3KooWANqay6T1xCrK3imfAEr3GpiGzxT6SWCEsRapo8tJWAZN",
		"/ip4/37.27.26.40/tcp/21888/p2p/12D3KooWPehhWgfdZG6x9zE7QdQxuAcfbQPAyBPjcXLfVYfU9FAd",
		"/dns/pactus-bootstrap.mflow.tech/tcp/21888/p2p/12D3KooWJD2PBAur6kXjX31dyrPzjSV947ppZ885QDKCXEvLm9Vb",
		"/ip4/5.161.222.154/tcp/21888/p2p/12D3KooWQN2VN1RWQ5eWq8rVexpBnwVSy7k34Rpep4R6V57XWYau",
		"/ip4/65.21.155.128/tcp/21888/p2p/12D3KooWPykSnzWpEasRijCgQ99vxAxWKL32zY24MmvtaXsHapSf",
		"/dns/pactus-bootstrap.teoviteovi.com/tcp/21888/p2p/12D3KooWMP5qWNHF1RfzhpAkBmQZzBfbaSJZCo3QF1fxs5CFXkv8",
		"/ip4/37.60.233.235/tcp/21888/p2p/12D3KooWNzbuBESYzbKGm7UyLkD22cHSgjRMUx4rjEKLBfVSZuY1",
		"/ip4/43.135.166.197/tcp/21888/p2p/12D3KooWEqyjMNzFe92HFZMZQd3YZ6zvztuPaC3a4wAUvPf4moyP",
		"/ip4/46.250.233.5/tcp/21888/p2p/12D3KooWH1kam9g1QFWmtWMLzP2MDY4ptpLTMPU8wSq3Q4k3rvnn",
		"/dns/pactus-bootstrap.codeblocklabs.com/tcp/21888/p2p/12D3KooWSdmPLqWbmv5x5pBiL5nHwKN5S2jhP96oXCKmoEmAhV4Z",
		"/ip4/84.247.184.95/tcp/21888/p2p/12D3KooWSjp5GFmCCsYkKcqicrghyHqfjdDCfLfgfqWmG5JTQSDu",
		"/dns/pactus-bootstrap.thd.io.vn/tcp/21888/p2p/12D3KooWDgsSSGJzihBJxqNJFqbGwbTEDXRoUBph1BCfwm11R6KY",
		"/ip4/37.27.25.245/tcp/21888/p2p/12D3KooWAmhn86b5igYjzFbaJddE57xT9UpZePFiEoc3jcuYHDvn",
		"/ip4/173.212.254.120/tcp/21888/p2p/12D3KooWJR19vjuBzuJTceCoy3LG8jggkHCBGb9j5SzdKtoPL7Yb",
		"/ip4/62.171.134.3/tcp/21888/p2p/12D3KooWSv5jsXMg1b4PBd2SCcPLj3Mc5KzzBYnpPkiS7JjuqhHF",
		"/ip4/109.123.240.5/tcp/21888/p2p/12D3KooWCT92yaS4JgrHHzeRPsJxaLzRdSEknGEN4HxQFfscoY8k",
		"/ip4/75.119.159.113/tcp/21888/p2p/12D3KooWL7nUsmd7JNrhrfbBjg6JxPFiZA5LXbdCDPnw5Xpdp4a9",
		"/ip4/161.97.145.100/tcp/21888/p2p/12D3KooWCjZ35Hu8Mt5GjZ6fJzJBZuUSsAppAzXcbMcXkswR4HrU",
		"/ip4/109.123.253.88/tcp/21888/p2p/12D3KooWFUJhDjYcAXgEnW1pTLRDWnH9CLgBPcezgU7pCpxMngt8",
		"/ip4/167.86.80.169/tcp/21888/p2p/12D3KooWDqAPjeRR41GYoQw9RQaSZdFx59k3SpkR29yAycpJUhfF",
		"/ip4/65.21.69.53/tcp/21888/p2p/12D3KooWQ1izUE5jcmMyjmxRdb2CmASe8Nfn81UL2G5rYjptC2RV",
		"/dns/pactus-bootstrap.validator.wiki/tcp/21888/p2p/12D3KooWBC5YQJdRAvupSXUPNYHiSctzZeTPNzen1JsRwj9aEyYZ",
		"/ip4/157.90.22.45/tcp/21888/p2p/12D3KooWLUYJ55DyJjSZvwPJDQNVYSBYqiMxJ9bFzRcVBgvSkN9L",
		"/ip4/65.21.152.153/tcp/21888/p2p/12D3KooWCirKBMmhmL6a4Ng5TVHQ4Leb6Wfj4M8s9raUzdFQn5aP",
		"/ip4/128.140.103.90/tcp/21888/p2p/12D3KooWN4eypkp61j12SMSkmrEfQv4RjnggmRfcEfom4mcKUpsR",
		"/ip4/185.216.75.142/tcp/21888/p2p/12D3KooWNjYFZpq5YwdUC6JgHfkZhUY7SyxTAQCHrdNwf2t1FXSn",
		"/ip4/85.190.246.251/tcp/21888/p2p/12D3KooWMTPKpr6EHMmZ2NSXkup4NaqwKnYmScHEsN3HpcqYgBZL",
		"/ip4/45.134.226.48/tcp/21888/p2p/12D3KooWBSP8AdiHkkvPunpbhfZE96YTY6om4hw35Xpe3bJFBUjh",
		"/ip4/154.38.187.40/tcp/21888/p2p/12D3KooWQyt7JCyhhnZ8jyT1vso4GGYVbpJhoXJy5KRpnqa8DmPg",
		"/ip4/194.163.164.10/tcp/21888/p2p/12D3KooWDRFCW5ksbnn6Uq7kx3HZt3EnrKjujqqs8QBQmA7TK5o6",
		"/ip4/167.86.116.61/tcp/21888/p2p/12D3KooWAmc69XaFh6T76BQQaFmSakYFgDyvzDLfVBgTzqmNtyXM",
		"/ip4/46.196.214.35/tcp/21888/p2p/12D3KooWJ3H7YAyemqxtToG6yRRzTfLdZD3eJk2xsQRaKxfdNhKf",
		"/ip4/42.117.247.180/tcp/21888/p2p/12D3KooWLVZSkgkFgrRNMfkCtULqtErYcNgjqgueeH944vDpDCPv",
		"/ip4/38.242.225.172/tcp/21888/p2p/12D3KooWDTP7CKGvbcg5WiuKB2Bt4acf2LWMdbUjMqU63dKKCVMB",
		"/ip4/167.86.121.11/tcp/21888/p2p/12D3KooWMpVXMQkr26ff3y8gC5z1h9RZNxTvW4WzZBML8GHHUoL5",
		"/ip4/194.163.189.255/tcp/21888/p2p/12D3KooWExbN79BrgDWjdLUd8TZaP2LvCDN8zJkMuPhRHkjKDcJS",
		"/ip4/37.60.245.47/tcp/21888/p2p/12D3KooWBH27VEPdGBYzq1vYciRgnDn3ozStRaUYd2vJWGgaKEUZ",
		"/ip4/84.247.172.23/tcp/21888/p2p/12D3KooWFJYk55mUD6FR9Gx2YGQ34PYziwTQhY1kZ3TMS6fGUz6g",
		"/ip4/95.216.222.135/tcp/21888/p2p/12D3KooWHicbZ3dBnR7JaVKr4U7WNUWyhDnUn2TwMVmPqW8Ftqop",
		"/ip4/65.109.143.78/tcp/21888/p2p/12D3KooWPnWgEFV32cvpfuv98oioZQJJKvRM46BEsFaoXYijvgvc",
		"/ip4/116.203.238.65/tcp/21888/p2p/12D3KooWKf32k9dLE4fA5tkh4KvvEAQDmHxHLWiahjpt2Hfx3d88",
		"/ip4/49.13.27.37/tcp/21888/p2p/12D3KooWPSh9KMa6tRfSewK3sUC4BDSPnXseVDGJxWqNZxKcXnGM",
		"/ip4/37.60.241.89/tcp/21888/p2p/12D3KooWMZtkk5HS69ukYYdfxD3CgBgTtKZg7kv43mCz1HMrDnQw",
		"/ip4/195.201.112.164/tcp/21888/p2p/12D3KooWKiYGT4mCapgcqLPSwbkJY1KgZMW5hdC8XJMQSU23tiSL",
		"/ip4/94.72.120.50/tcp/21888/p2p/12D3KooWCJX9aLnzmEFtEasvbbaFqHLTboFVFptZxeFnbnSxpb2m",
		"/ip4/95.111.242.225/tcp/21888/p2p/12D3KooWRNTZpo1XNWKQx2Z5HEu5RnqKeNDV5y4qMKX5f4aLtbv9",
		"/dns/pactus-bootstrap1.dezh.tech/tcp/21888/p2p/12D3KooWK1z7QAskVrQd98r98UMSPwTLr4out8B9NkQTrCdZPCZx",
	}

	BootstrapPeersBitcoinDNSSeeds = []string{
		"seed.bitcoin.sipa.be",
		"dnsseed.bluematt.me",
		"dnsseed.bitcoin.dashjr.org",
		"seed.bitcoinstats.com",
		"seed.bitnodes.io",
		"bitseed.xf2.org",
	}
	// BootstrapPeersPortalMainnet extracted from:
	//	https://github.com/status-im/portal-mainnet/blob/master/config/bootstrap_nodes.txt
	//  https://github.com/ethereum/portal-network-specs/blob/master/bootnodes.md
	BootstrapPeersPortalMainnet = []string{
		// Fluffy bootstrap nodes
		// nimbus-fluffy-mainnet-master-01@metal-01.ih-eu-mda1.nimbus.fluffy
		"enr:-Ia4QLBxlH0Y8hGPQ1IRF5EStZbZvCPHQ2OjaJkuFMz0NRoZIuO2dLP0L-W_8ZmgnVx5SwvxYCXmX7zrHYv0FeHFFR0TY2aCaWSCdjSCaXCEwiErIIlzZWNwMjU2azGhAnnTykipGqyOy-ZRB9ga9pQVPF-wQs-yj_rYUoOqXEjbg3VkcIIjjA",
		// nimbus-fluffy-mainnet-master-02@metal-01.ih-eu-mda1.nimbus.fluffy
		"enr:-Ia4QM4amOkJf5z84Lv5Fl0RgWeSSDUekwnOPRn6XA1eMWgrHwWmn_gJGtOeuVfuX7ywGuPMRwb0odqQ9N_w_2Qc53gTY2aCaWSCdjSCaXCEwiErIYlzZWNwMjU2azGhAzaQEdPmz9SHiCw2I5yVAO8sriQ-mhC5yB7ea1u4u5QZg3VkcIIjjA",
		// nimbus-fluffy-mainnet-master-01@metal-02.ih-eu-mda1.nimbus.fluffy
		"enr:-Ia4QKVuHjNafkYuvhU7yCvSarNIVXquzJ8QOp5YbWJRIJw_EDVOIMNJ_fInfYoAvlRCHEx9LUQpYpqJa04pUDU21uoTY2aCaWSCdjSCaXCEwiErQIlzZWNwMjU2azGhA47eAW5oIDJAqxxqI0sL0d8ttXMV0h6sRIWU4ZwS4pYfg3VkcIIjjA",
		// nimbus-fluffy-mainnet-master-02@metal-02.ih-eu-mda1.nimbus.fluffy
		"enr:-Ia4QIU9U3zrP2DM7sfpgLJbbYpg12sWeXNeYcpKN49-6fhRCng0IUoVRI2E51mN-2eKJ4tbTimxNLaAnbA7r7fxVjcTY2aCaWSCdjSCaXCEwiErQYlzZWNwMjU2azGhAxOroJ3HceYvdD2yK1q9w8c9tgrISJso8q_JXI6U0Xwng3VkcIIjjA",

		// Trin bootstrap nodes
		"enr:-Jy4QIs2pCyiKna9YWnAF0zgf7bT0GzlAGoF8MEKFJOExmtofBIqzm71zDvmzRiiLkxaEJcs_Amr7XIhLI74k1rtlXICY5Z0IDAuMS4xLWFscGhhLjEtMTEwZjUwgmlkgnY0gmlwhKEjVaWJc2VjcDI1NmsxoQLSC_nhF1iRwsCw0n3J4jRjqoaRxtKgsEe5a-Dz7y0JloN1ZHCCIyg",
		"enr:-Jy4QKSLYMpku9F0Ebk84zhIhwTkmn80UnYvE4Z4sOcLukASIcofrGdXVLAUPVHh8oPCfnEOZm1W1gcAxB9kV2FJywkCY5Z0IDAuMS4xLWFscGhhLjEtMTEwZjUwgmlkgnY0gmlwhJO2oc6Jc2VjcDI1NmsxoQLMSGVlxXL62N3sPtaV-n_TbZFCEM5AR7RDyIwOadbQK4N1ZHCCIyg",
		"enr:-Jy4QH4_H4cW--ejWDl_W7ngXw2m31MM2GT8_1ZgECnfWxMzZTiZKvHDgkmwUS_l2aqHHU54Q7hcFSPz6VGzkUjOqkcCY5Z0IDAuMS4xLWFscGhhLjEtMTEwZjUwgmlkgnY0gmlwhJ31OTWJc2VjcDI1NmsxoQPC0eRkjRajDiETr_DRa5N5VJRm-ttCWDoO1QAMMCg5pIN1ZHCCIyg",

		// Ultralight bootstrap nodes
		"enr:-IS4QFV_wTNknw7qiCGAbHf6LxB-xPQCktyrCEZX-b-7PikMOIKkBg-frHRBkfwhI3XaYo_T-HxBYmOOQGNwThkBBHYDgmlkgnY0gmlwhKRc9_OJc2VjcDI1NmsxoQKHPt5CQ0D66ueTtSUqwGjfhscU_LiwS28QvJ0GgJFd-YN1ZHCCE4k",
		"enr:-IS4QDpUz2hQBNt0DECFm8Zy58Hi59PF_7sw780X3qA0vzJEB2IEd5RtVdPUYZUbeg4f0LMradgwpyIhYUeSxz2Tfa8DgmlkgnY0gmlwhKRc9_OJc2VjcDI1NmsxoQJd4NAVKOXfbdxyjSOUJzmA4rjtg43EDeEJu1f8YRhb_4N1ZHCCE4o",
		"enr:-IS4QGG6moBhLW1oXz84NaKEHaRcim64qzFn1hAG80yQyVGNLoKqzJe887kEjthr7rJCNlt6vdVMKMNoUC9OCeNK-EMDgmlkgnY0gmlwhKRc9-KJc2VjcDI1NmsxoQLJhXByb3LmxHQaqgLDtIGUmpANXaBbFw3ybZWzGqb9-IN1ZHCCE4k",
		"enr:-IS4QA5hpJikeDFf1DD1_Le6_ylgrLGpdwn3SRaneGu9hY2HUI7peHep0f28UUMzbC0PvlWjN8zSfnqMG07WVcCyBhADgmlkgnY0gmlwhKRc9-KJc2VjcDI1NmsxoQJMpHmGj1xSP1O-Mffk_jYIHVcg6tY5_CjmWVg1gJEsPIN1ZHCCE4o",
	}

	// BootstrapPeersDria extracted from:
	//   https://github.com/firstbatchxyz/dkn-compute-node/blob/b4f21e19197df48d4224c466109df3ff7c85ff50/p2p/src/network.rs#L47
	BootstrapPeersDria = []string{
		"/ip4/44.206.245.139/tcp/4001/p2p/16Uiu2HAm4q3LZU2T9kgjKK4ysy6KZYKLq8KiXQyae4RHdF7uqSt4",
		"/ip4/18.234.39.91/tcp/4001/p2p/16Uiu2HAmJqegPzwuGKWzmb5m3RdSUJ7NhEGWB5jNCd3ca9zdQ9dU",
		"/ip4/54.242.44.217/tcp/4001/p2p/16Uiu2HAmR2sAoh9F8jT9AZup9y79Mi6NEFVUbwRvahqtWamfabkz",
		"/ip4/52.201.242.227/tcp/4001/p2p/16Uiu2HAmFEUCy1s1gjyHfc8jey4Wd9i5bSDnyFDbWTnbrF2J3KFb",
	}

	// BootstrapPeersWakuStatus extracted from:
	//	https://github.com/status-im/status-mobile/blob/develop/resources/config/fleets.json
	BootstrapPeersWakuStatus = []string{
		"enr:-QEKuED9AJm2HGgrRpVaJY2nj68ao_QiPeUT43sK-aRM7sMJ6R4G11OSDOwnvVacgN1sTw-K7soC5dzHDFZgZkHU0u-XAYJpZIJ2NIJpcISnYxMvim11bHRpYWRkcnO4WgAqNiVib290LTAxLmRvLWFtczMuc3RhdHVzLnByb2Quc3RhdHVzLmltBnZfACw2JWJvb3QtMDEuZG8tYW1zMy5zdGF0dXMucHJvZC5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEC3rRtFQSgc24uWewzXaxTY8hDAHB8sgnxr9k8Rjb5GeSDdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
		"enr:-QEcuED7ww5vo2rKc1pyBp7fubBUH-8STHEZHo7InjVjLblEVyDGkjdTI9VdqmYQOn95vuQH-Htku17WSTzEufx-Wg4mAYJpZIJ2NIJpcIQihw1Xim11bHRpYWRkcnO4bAAzNi5ib290LTAxLmdjLXVzLWNlbnRyYWwxLWEuc3RhdHVzLnByb2Quc3RhdHVzLmltBnZfADU2LmJvb3QtMDEuZ2MtdXMtY2VudHJhbDEtYS5zdGF0dXMucHJvZC5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaECxjqgDQ0WyRSOilYU32DA5k_XNlDis3m1VdXkK9xM6kODdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
		"enr:-QEcuEAoShWGyN66wwusE3Ri8hXBaIkoHZHybUB8cCPv5v3ypEf9OCg4cfslJxZFANl90s-jmMOugLUyBx4EfOBNJ6_VAYJpZIJ2NIJpcIQI2hdMim11bHRpYWRkcnO4bAAzNi5ib290LTAxLmFjLWNuLWhvbmdrb25nLWMuc3RhdHVzLnByb2Quc3RhdHVzLmltBnZfADU2LmJvb3QtMDEuYWMtY24taG9uZ2tvbmctYy5zdGF0dXMucHJvZC5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDP7CbRk-YKJwOFFM4Z9ney0GPc7WPJaCwGkpNRyla7mCDdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
		"enr:-QEQuEBuiQgFlJNcv255042zwyl4pOBOivakX8N30Dr9vaaEU2q8-7N4GVY4Hk87iEKELjlIXTpE9Wj6EQq1lrBuc7ayAYJpZIJ2NIJpcISPxvrpim11bHRpYWRkcnO4YAAtNihib290LTAxLmRvLWFtczMuc3RhdHVzLnN0YWdpbmcuc3RhdHVzLmltBnZfAC82KGJvb3QtMDEuZG8tYW1zMy5zdGF0dXMuc3RhZ2luZy5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDq-yGgpuoUG6NKkbIDRmrMiT-bEVzFlpWLEK_rF3yKUaDdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
		"enr:-QEiuED2UusuHo1d6WN2-tHjtj0T0gdnsOh7aRZnFF6OEYLDbyxOtQo2_4dFUHhc9xm5SHNrWJJq8X7FRsxc4VCMGjjbAYJpZIJ2NIJpcIRoxQVgim11bHRpYWRkcnO4cgA2NjFib290LTAxLmdjLXVzLWNlbnRyYWwxLWEuc3RhdHVzLnN0YWdpbmcuc3RhdHVzLmltBnZfADg2MWJvb3QtMDEuZ2MtdXMtY2VudHJhbDEtYS5zdGF0dXMuc3RhZ2luZy5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDNAvlGjekD1YV4WpmjwArGAH2g9kHFJnMRfgUhcIkoA2DdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
		"enr:-QEiuECJPv2vL00Jp5sTEMAFyW7qXkK2cFgphlU_G8-FJuJqoW_D5aWIy3ylGdv2K8DkiG7PWgng4Ql_VI7Qc2RhBdwfAYJpZIJ2NIJpcIQvTKi6im11bHRpYWRkcnO4cgA2NjFib290LTAxLmFjLWNuLWhvbmdrb25nLWMuc3RhdHVzLnN0YWdpbmcuc3RhdHVzLmltBnZfADg2MWJvb3QtMDEuYWMtY24taG9uZ2tvbmctYy5zdGF0dXMuc3RhZ2luZy5zdGF0dXMuaW0GAbveA4Jyc40AEAUAAQAgAEAAgAEAiXNlY3AyNTZrMaEDkbgV7oqPNmFtX5FzSPi9WH8kkmrPB1R3n9xRXge91M-DdGNwgnZfg3VkcIIjKIV3YWt1Mg0",
	}

	// BootstrapPeersWakuTWN extracted from:
	//  fork from chaitanyaprem: https://github.com/dennis-tra/nebula/commit/acb8cc72958030c2be84b5f2cd90dc1dfeb327c6
	BootstrapPeersWakuTWN = []string{
		"enr:-QESuED0qW1BCmF-oH_ARGPr97Nv767bl_43uoy70vrbah3EaCAdK3Q0iRQ6wkSTTpdrg_dU_NC2ydO8leSlRpBX4pxiAYJpZIJ2NIJpcIRA4VDAim11bHRpYWRkcnO4XAArNiZub2RlLTAxLmRvLWFtczMud2FrdS5zYW5kYm94LnN0YXR1cy5pbQZ2XwAtNiZub2RlLTAxLmRvLWFtczMud2FrdS5zYW5kYm94LnN0YXR1cy5pbQYfQN4DgnJzkwABCAAAAAEAAgADAAQABQAGAAeJc2VjcDI1NmsxoQOTd-h5owwj-cx7xrmbvQKU8CV3Fomfdvcv1MBc-67T5oN0Y3CCdl-DdWRwgiMohXdha3UyDw",
		"enr:-QEkuED9X80QF_jcN9gA2ZRhhmwVEeJnsg_Hyg7IFCTYnZD0BDI7a8HArE61NhJZFwygpHCWkgwSt2vqiABXkBxzIqZBAYJpZIJ2NIJpcIQiQlleim11bHRpYWRkcnO4bgA0Ni9ub2RlLTAxLmdjLXVzLWNlbnRyYWwxLWEud2FrdS5zYW5kYm94LnN0YXR1cy5pbQZ2XwA2Ni9ub2RlLTAxLmdjLXVzLWNlbnRyYWwxLWEud2FrdS5zYW5kYm94LnN0YXR1cy5pbQYfQN4DgnJzkwABCAAAAAEAAgADAAQABQAGAAeJc2VjcDI1NmsxoQPFAS8zz2cg1QQhxMaK8CzkGQ5wdHvPJcrgLzJGOiHpwYN0Y3CCdl-DdWRwgiMohXdha3UyDw",
		"enr:-QEkuEBfEzJm_kigJ2HoSS_RBFJYhKHocGdkhhBr6jSUAWjLdFPp6Pj1l4yiTQp7TGHyu1kC6FyaU573VN8klLsEm-XuAYJpZIJ2NIJpcIQI2SVcim11bHRpYWRkcnO4bgA0Ni9ub2RlLTAxLmFjLWNuLWhvbmdrb25nLWMud2FrdS5zYW5kYm94LnN0YXR1cy5pbQZ2XwA2Ni9ub2RlLTAxLmFjLWNuLWhvbmdrb25nLWMud2FrdS5zYW5kYm94LnN0YXR1cy5pbQYfQN4DgnJzkwABCAAAAAEAAgADAAQABQAGAAeJc2VjcDI1NmsxoQOwsS69tgD7u1K50r5-qG5hweuTwa0W26aYPnvivpNlrYN0Y3CCdl-DdWRwgiMohXdha3UyDw",
	}

	// BootstrapPeersGnosis extracted from:
	//  https://github.com/gnosischain/configs/blob/main/mainnet/bootnodes.yaml
	BootstrapPeersGnosis = []string{
		"enr:-Ly4QIAhiTHk6JdVhCdiLwT83wAolUFo5J4nI5HrF7-zJO_QEw3cmEGxC1jvqNNUN64Vu-xxqDKSM528vKRNCehZAfEBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCCS-QxAgAAZP__________gmlkgnY0gmlwhEFtZ5SJc2VjcDI1NmsxoQJwgL5C-30E8RJmW8gCb7sfwWvvfre7wGcCeV4X1G2wJYhzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA",
		"enr:-Ly4QDhEjlkf8fwO5uWAadexy88GXZneTuUCIPHhv98v8ZfXMtC0S1S_8soiT0CMEgoeLe9Db01dtkFQUnA9YcnYC_8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCCS-QxAgAAZP__________gmlkgnY0gmlwhEFtZ5WJc2VjcDI1NmsxoQMRSho89q2GKx_l2FZhR1RmnSiQr6o_9hfXfQUuW6bjMohzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA",
		"enr:-Ly4QLKgv5M2D4DYJgo6s4NG_K4zu4sk5HOLCfGCdtgoezsbfRbfGpQ4iSd31M88ec3DHA5FWVbkgIas9EaJeXia0nwBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCCS-QxAgAAZP__________gmlkgnY0gmlwhI1eYRaJc2VjcDI1NmsxoQLpK_A47iNBkVjka9Mde1F-Kie-R0sq97MCNKCxt2HwOIhzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA",
		"enr:-Ly4QF_0qvji6xqXrhQEhwJR1W9h5dXV7ZjVCN_NlosKxcgZW6emAfB_KXxEiPgKr_-CZG8CWvTiojEohG1ewF7P368Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCCS-QxAgAAZP__________gmlkgnY0gmlwhI1eYUqJc2VjcDI1NmsxoQIpNRUT6llrXqEbjkAodsZOyWv8fxQkyQtSvH4sg2D7n4hzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA",
		"enr:-Ly4QCD5D99p36WafgTSxB6kY7D2V1ca71C49J4VWI2c8UZCCPYBvNRWiv0-HxOcbpuUdwPVhyWQCYm1yq2ZH0ukCbQBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCCS-QxAgAAZP__________gmlkgnY0gmlwhI1eYVSJc2VjcDI1NmsxoQJJMSV8iSZ8zvkgbi8cjIGEUVJeekLqT0LQha_co-siT4hzeW5jbmV0cwCDdGNwgiMog3VkcIIjKA",
		"enr:-KK4QKXJq1QOVWuJAGige4uaT8LRPQGCVRf3lH3pxjaVScMRUfFW1eiiaz8RwOAYvw33D4EX-uASGJ5QVqVCqwccxa-Bi4RldGgykCGm-DYDAABk__________-CaWSCdjSCaXCEM0QnzolzZWNwMjU2azGhAhNvrRkpuK4MWTf3WqiOXSOePL8Zc-wKVpZ9FQx_BDadg3RjcIIjKIN1ZHCCIyg",
		"enr:-LO4QO87Rn2ejN3SZdXkx7kv8m11EZ3KWWqoIN5oXwQ7iXR9CVGd1dmSyWxOL1PGsdIqeMf66OZj4QGEJckSi6okCdWBpIdhdHRuZXRziAAAAABgAAAAhGV0aDKQPr_UhAQAAGT__________4JpZIJ2NIJpcIQj0iX1iXNlY3AyNTZrMaEDd-_eqFlWWJrUfEp8RhKT9NxdYaZoLHvsp3bbejPyOoeDdGNwgiMog3VkcIIjKA",
		"enr:-LK4QIJUAxX9uNgW4ACkq8AixjnSTcs9sClbEtWRq9F8Uy9OEExsr4ecpBTYpxX66cMk6pUHejCSX3wZkK2pOCCHWHEBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpA-v9SEBAAAZP__________gmlkgnY0gmlwhCPSnDuJc2VjcDI1NmsxoQNuaAjFE-ANkH3pbeBdPiEIwjR5kxFuKaBWxHkqFuPz5IN0Y3CCIyiDdWRwgiMo",
	}

	BootstrapPeersAlgoTestnet = []string{
		"/ip4/167.99.134.197/tcp/4161/p2p/12D3KooWEGB9Bt6X6CTBypQDQmcrkU7kUVeTdxXbYRKPatWkn3bt",
	}

	BootstrapPeersAlgoMainnet = []string{
		"/ip4/64.227.112.163/tcp/4161/p2p/12D3KooWAUFhQq394HsYTw18bqG3dwsfrGScZceCK7QsieakVMst",
	}

	BootstrapPeersAlgoSuppranet = []string{
		"/ip4/64.227.112.163/tcp/4161/p2p/12D3KooWAUFhQq394HsYTw18bqG3dwsfrGScZceCK7QsieakVMst",
	}
)
