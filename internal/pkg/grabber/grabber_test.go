package grabber

import (
	"testing"
)

func BenchmarkTranslit(b *testing.B) {
	testString := `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	Id est enim, de quo quaerimus. Laboro autem non sine causa; Bork Nam quid possumus facere melius?
	Bork Stoicos roga. Nunc omni virtuti vitium contrario nomine opponitur.
	Quod equidem non reprehendo; Cave putes quicquam esse verius. Restinguet citius, si ardentem acceperit.
	Recte dicis; Sed haec omittamus; Tum Torquatus: Prorsus, inquit, assentior; Laboro autem non sine causa;
	Facile est hoc cernere in primis puerorum aetatulis. Quae cum dixisset, finem ille.
	Duo Reges: constructio interrete. Satis est ad hoc responsum.
	Omnes enim iucundum motum, quo sensus hilaretur. Cave putes quicquam esse verius.
	Polycratem Samium felicem appellabant. Non semper, inquam;
	Tum ille: Ain tandem? Quid sequatur, quid repugnet, vident.`

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		translit(testString)
	}
}

// Test translit function
func TestTranslit(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc: "successful english translit",
			input: `Sed quid attinet de rebus tam apertis plura requirere? Duo Reges: constructio interrete.
			Negare non possum. Hoc loco tenere se Triarius non potuit. Itaque contra est, ac dicitis;
			In schola desinis. `,
			want: "sed-quid-attinet-de-rebus-tam-apertis-plura-requirere-duo-reges-constructio-interrete-negare-non-possum-hoc-loco-tenere-se-triarius-non-potuit-itaque-contra-est-ac-dicitis-in-schola-desinis",
		},
		{
			desc: "successful russian translit",
			input: `Пример строки для тестирования функции Транслитерации!
			Здесь есть и русские буквы (Ёё), и английские (ABCD), и знаки препинания.`,
			want: "primer-stroki-dlia-testirovaniia-funkcii-transliteracii-zdes-est-i-russkie-bukvi-ee-i-angliiskie-abcd-i-znaki-prepinaniia",
		},
		{
			desc: "trimm - ",
			input: "-----input----вход-----",
			want: "input-vhod",
		},
		{
			desc: "trimm space",
			input: "   input    ",
			want: "input",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := translit(tC.input)
			if res != tC.want {
				t.Errorf("\n%s: \n\twant:\n\t\"%s\" \n\tget: \n\t\"%s\"", tC.desc, tC.want, res)
			}
		})
	}
}
