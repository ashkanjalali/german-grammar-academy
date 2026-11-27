package academy

import (
	"encoding/json"

	"gorm.io/gorm"
)

type lessonSeed struct {
	level, title, slug, description, explanation string
	examples []map[string]string
}

func MigrateAndSeed(db *gorm.DB) error {
	if err := db.AutoMigrate(&Lesson{}, &Question{}, &Attempt{}, &Progress{}); err != nil { return err }
	var count int64
	db.Model(&Lesson{}).Count(&count)
	if count > 0 { return nil }

	seeds := []lessonSeed{
		{"A1","Definite and indefinite articles","a1-articles","Learn der, die, das and ein/eine.","German nouns have grammatical gender. Learn each noun together with its article. Use der for masculine, die for feminine and das for neuter nouns.",[]map[string]string{{"de":"der Tisch","en":"the table (masculine)"},{"de":"die Lampe","en":"the lamp (feminine)"},{"de":"das Buch","en":"the book (neuter)"}}},
		{"A1","Present tense: regular verbs","a1-present","Conjugate common verbs in the Präsens.","Remove -en from the infinitive and add the personal ending: ich -e, du -st, er/sie/es -t, wir -en, ihr -t, sie/Sie -en.",[]map[string]string{{"de":"Ich lerne Deutsch.","en":"I am learning German."},{"de":"Du wohnst in Berlin.","en":"You live in Berlin."}}},
		{"A2","The dative case","a2-dative","Use the dative for indirect objects and after dative prepositions.","The dative often marks the recipient. Masculine and neuter articles become dem, feminine becomes der, and plural becomes den (+n where possible).",[]map[string]string{{"de":"Ich helfe dem Mann.","en":"I help the man."},{"de":"Das Buch gehört der Frau.","en":"The book belongs to the woman."}}},
		{"A2","Perfekt with haben and sein","a2-perfekt","Talk about completed past events.","The conversational past is formed with haben or sein plus a past participle. Many movement or change-of-state verbs use sein.",[]map[string]string{{"de":"Ich habe Deutsch gelernt.","en":"I learned German."},{"de":"Wir sind nach Hause gegangen.","en":"We went home."}}},
		{"B1","Subordinate clauses with weil and dass","b1-subclauses","Put the conjugated verb at the end of a subordinate clause.","After subordinating conjunctions such as weil, dass and obwohl, the finite verb moves to the end of its clause.",[]map[string]string{{"de":"Ich lerne, weil ich eine Prüfung habe.","en":"I am studying because I have an exam."},{"de":"Sie sagt, dass sie müde ist.","en":"She says that she is tired."}}},
		{"B1","Konjunktiv II for polite requests","b1-konjunktiv","Make polite requests and describe hypothetical situations.","Konjunktiv II forms such as könnte, hätte and wäre soften requests and express unreal or hypothetical situations.",[]map[string]string{{"de":"Könnten Sie mir helfen?","en":"Could you help me?"},{"de":"Wenn ich Zeit hätte, würde ich reisen.","en":"If I had time, I would travel."}}},
		{"B2","Passive voice","b2-passive","Focus on an action rather than its agent.","The process passive uses werden + past participle. The state passive uses sein + past participle.",[]map[string]string{{"de":"Der Brief wird geschrieben.","en":"The letter is being written."},{"de":"Die Tür ist geschlossen.","en":"The door is closed."}}},
		{"B2","Relative clauses","b2-relative","Add precise information about a noun.","Relative pronouns agree in gender and number with the antecedent; their case depends on their role inside the relative clause.",[]map[string]string{{"de":"Das ist der Mann, dem ich geholfen habe.","en":"That is the man whom I helped."},{"de":"Die Frau, die dort steht, ist Ärztin.","en":"The woman who is standing there is a doctor."}}},
		{"C1","Nominalisation","c1-nominalisierung","Use noun phrases for concise formal writing.","German formal and academic registers often turn verbs and adjectives into nouns. Pay attention to capitalization and case government.",[]map[string]string{{"de":"Nachdem er angekommen war, begann die Sitzung.","en":"After he had arrived, the meeting began."},{"de":"Nach seiner Ankunft begann die Sitzung.","en":"After his arrival, the meeting began."}}},
		{"C1","Complex conjunctions and connectors","c1-konnektoren","Express contrast, concession and consequence precisely.","Connectors such as insofern als, wenngleich and nichtsdestotrotz express nuanced logical relationships and influence clause structure.",[]map[string]string{{"de":"Wenngleich die Lage schwierig ist, bleiben wir optimistisch.","en":"Although the situation is difficult, we remain optimistic."}}},
		{"C2","Register and syntactic nuance","c2-register","Choose structures appropriate to context and register.","At C2, grammatical accuracy includes control of information structure, emphasis, ellipsis and register. A grammatically possible sentence may still be pragmatically unsuitable.",[]map[string]string{{"de":"Was die Umsetzung betrifft, bestehen weiterhin offene Fragen.","en":"As far as implementation is concerned, questions remain."}}},
		{"C2","Advanced participial constructions","c2-partizipial","Use concise written participial modifiers accurately.","Expanded participial attributes can compress relative clauses in formal writing. Ensure the modifier clearly attaches to the intended noun.",[]map[string]string{{"de":"Die gestern veröffentlichte Studie sorgt für Diskussionen.","en":"The study published yesterday is causing debate."}}},
	}
	for i, s := range seeds {
		ex, _ := json.Marshal(s.examples)
		l := Lesson{Level:s.level, Title:s.title, Slug:s.slug, Description:s.description, Explanation:s.explanation, Examples:string(ex), Order:i+1}
		if err := db.Create(&l).Error; err != nil { return err }
	}
	var lessons []Lesson
	db.Find(&lessons)
	questions := []Question{
		{Level:"A1", LessonID:lessons[0].ID, Prompt:"___ Buch ist interessant.", Options:`["Der","Die","Das","Dem"]`, CorrectAnswer:2, Explanation:"Buch is neuter, so the nominative definite article is das."},
		{Level:"A1", LessonID:lessons[1].ID, Prompt:"Ich ___ Deutsch.", Options:`["lernst","lerne","lernt","lernen"]`, CorrectAnswer:1, Explanation:"The first-person singular ending is -e: ich lerne."},
		{Level:"A2", LessonID:lessons[2].ID, Prompt:"Ich helfe ___ Frau.", Options:`["die","der","den","das"]`, CorrectAnswer:1, Explanation:"Helfen takes the dative; feminine dative is der."},
		{Level:"A2", LessonID:lessons[3].ID, Prompt:"Gestern ___ ich nach Hause gegangen.", Options:`["habe","bin","hat","werde"]`, CorrectAnswer:1, Explanation:"Gehen forms the Perfekt with sein."},
		{Level:"B1", LessonID:lessons[4].ID, Prompt:"Ich bleibe zu Hause, weil ich krank ___.", Options:`["bin","ist","sein","warst"]`, CorrectAnswer:0, Explanation:"The finite verb goes to the end of the weil-clause."},
		{Level:"B1", LessonID:lessons[5].ID, Prompt:"___ Sie das bitte wiederholen?", Options:`["Könnten","Könntest","Kann","Konntet"]`, CorrectAnswer:0, Explanation:"Könnten Sie is a polite formal request."},
		{Level:"B2", LessonID:lessons[6].ID, Prompt:"Der Bericht ___ morgen veröffentlicht.", Options:`["hat","wird","ist","wurde"]`, CorrectAnswer:1, Explanation:"Present passive: werden + past participle."},
		{Level:"B2", LessonID:lessons[7].ID, Prompt:"Das ist die Kollegin, ___ ich gestern begegnet bin.", Options:`["die","der","den","das"]`, CorrectAnswer:1, Explanation:"Begegnen takes the dative; feminine dative relative pronoun is der."},
		{Level:"C1", LessonID:lessons[8].ID, Prompt:"Choose the nominalised form: 'nachdem er angekommen war'.", Options:`["nachdem Ankommen","nach seiner Ankunft","nachdem er Ankunft","nach ankommen"]`, CorrectAnswer:1, Explanation:"Nach seiner Ankunft is the nominalised construction."},
		{Level:"C2", LessonID:lessons[10].ID, Prompt:"Which sentence uses a participial attribute correctly?", Options:`["Die gestern veröffentlichte Studie ist neu.","Die gestern veröffentlichen Studie ist neu.","Die gestern veröffentlichter Studie ist neu.","Die gestern veröffentlichend Studie ist neu."]`, CorrectAnswer:0, Explanation:"Veröffentlichte agrees with feminine singular Studie in the nominative."},
	}
	for _, q := range questions { if err := db.Create(&q).Error; err != nil { return err } }
	return nil
}
