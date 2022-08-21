package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.SeekBar;
import android.widget.TextView;
import android.widget.Toast;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.tasks.BetTask;

public class BetActivity extends AppCompatActivity {

    private String token;
    private String leagueId;
    private String userId;
    private String homeId;
    private String visitorsId;
    private int home;
    private int visitors;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_bet);

        Intent intent = getIntent();
        token = intent.getStringExtra("token");
        leagueId = intent.getStringExtra("leagueId");
        userId = intent.getStringExtra("userId");
        homeId = intent.getStringExtra("homeId");
        visitorsId = intent.getStringExtra("visitorsId");
        home = intent.getIntExtra("home", 0);
        visitors = intent.getIntExtra("visitors", 0);

        TextView tvHomeId = findViewById(R.id.tvHomeId);
        TextView tvVisitorsId = findViewById(R.id.tvVisitorsId);
        EditText etHome = findViewById(R.id.etHome);
        EditText etVisitors = findViewById(R.id.etVisitors);
        SeekBar sbHome = findViewById(R.id.sbHome);
        SeekBar sbVisitors = findViewById(R.id.sbVisitors);

        tvHomeId.setText(homeId);
        tvVisitorsId.setText(visitorsId);
        etHome.setText(String.valueOf(home));
        etVisitors.setText(String.valueOf(visitors));
        sbHome.setProgress(home);
        sbVisitors.setProgress(visitors);

        sbHome.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            @Override
            public void onProgressChanged(SeekBar seekBar, int i, boolean b) {
                etHome.setText(String.valueOf(i));
            }

            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {

            }

            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {

            }
        });

        sbVisitors.setOnSeekBarChangeListener(new SeekBar.OnSeekBarChangeListener() {
            @Override
            public void onProgressChanged(SeekBar seekBar, int i, boolean b) {
                etVisitors.setText(String.valueOf(i));
            }

            @Override
            public void onStartTrackingTouch(SeekBar seekBar) {

            }

            @Override
            public void onStopTrackingTouch(SeekBar seekBar) {

            }
        });

    }

    public void confirm(View v) {
        SeekBar sbHome = findViewById(R.id.sbHome);
        home = sbHome.getProgress();
        SeekBar sbVisitor = findViewById(R.id.sbVisitors);
        visitors = sbVisitor.getProgress();
        new BetTask(this, leagueId, userId, homeId + "-" + visitorsId, home, visitors, token).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response == null) {
            Toast.makeText(this, "Sem resposta so servidor. Tente novamente mais tarde.", Toast.LENGTH_LONG).show();
            return;
        }
        if (response.code() == 204) {
            Toast.makeText(this, "Aposta registrada com sucesso", Toast.LENGTH_LONG).show();
            setResult(RESULT_OK);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
        finish();
    }
}