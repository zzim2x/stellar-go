package scp

type testDriver struct {
	quorumSets map[Hash]*QuorumSet
}

var _, _ Driver = &testDriver{}, (*testDriver)(nil)

func newTestDriver() *testDriver {
	return &testDriver{
		quorumSets: make(map[Hash]*QuorumSet),
	}
}

func (o *testDriver) VerifyEnvelope(envelope Envelope) bool {
	return false
}

func (o *testDriver) EmitEnvelope(envelope Envelope) {
}

func (o *testDriver) SignEnvelope(envelope *Envelope) {
}

func (o *testDriver) GetQuorumSet(hash Hash) *QuorumSet {
	return o.quorumSets[hash]
}

func (o *testDriver) storeQuorumSet(quorumSet *QuorumSet) {
	o.quorumSets[quorumSet.Hash()] = quorumSet
}

func (o *testDriver) ValidateValue(slotId uint64, value Value, nomination bool) ValidationLevel {
	return 0
}

func newNomination(slotIndex uint64, secretKey SecretKey, quorumSetHash Hash) Envelope {
	return makeEnvelope(secretKey, Statement{
		SlotIndex:     slotIndex,
		NodeId:        secretKey.PublicKey,
		StatementType: StatementTypeNomination,
		Nomination: &Nomination{
			QuorumSetHash: quorumSetHash,
			Votes:         []Value{},
			Accepted:      []Value{},
		},
	})
}

func makeEnvelope(secretKey SecretKey, statement Statement) Envelope {
	return Envelope{
		Statement: statement,
	}
}