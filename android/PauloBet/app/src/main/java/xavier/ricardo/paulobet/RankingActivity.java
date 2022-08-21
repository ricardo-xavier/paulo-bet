package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ListView;
import android.widget.Toast;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.adapters.RankingAdapter;
import xavier.ricardo.paulobet.model.Ranking;
import xavier.ricardo.paulobet.model.RankingResponse;
import xavier.ricardo.paulobet.tasks.RankingTask;

public class RankingActivity extends AppCompatActivity {
    private RankingActivity context;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        context = this;
        setContentView(R.layout.activity_scores);
        Intent intent = getIntent();
        String league = intent.getStringExtra("league");
        String user = intent.getStringExtra("user");
        String token = intent.getStringExtra("token");

        ListView lvScores = findViewById(R.id.lvScores);
        lvScores.setOnItemClickListener(new AdapterView.OnItemClickListener() {
            @Override
            public void onItemClick(AdapterView<?> adapterView, View view, int i, long l) {
                Ranking ranking = (Ranking) adapterView.getAdapter().getItem(i);
                Intent intent = new Intent(context, GetBetsActivity.class);
                intent.putExtra("league", league);
                intent.putExtra("user", ranking.getUserId());
                intent.putExtra("login", user);
                intent.putExtra("token", token);
                startActivity(intent);
                System.out.println(i);
            }
        });

        new RankingTask(this, league, user, token).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response == null) {
            Toast.makeText(this, "Sem resposta so servidor. Tente novamente mais tarde.", Toast.LENGTH_LONG).show();
            return;
        }
        if (response.code() == 200) {
            String body = response.body().string();
            Gson gson = new Gson();
            RankingResponse rankingResponse = gson.fromJson(body, RankingResponse.class);
            ListView lvScores = findViewById(R.id.lvScores);
            RankingAdapter adapter = new RankingAdapter(this, rankingResponse.getRanking());
            lvScores.setAdapter(adapter);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
    }
}